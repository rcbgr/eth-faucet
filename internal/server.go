package internal

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"net"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	faucetpb "github.com/rauljordan/eth-faucet/proto/faucet"
	gateway "github.com/rauljordan/minimal-grpc-gateway"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	log = logrus.WithField("prefix", "server")
)

// Config for the faucet server.
type Config struct {
	GrpcPort          int      `mapstructure:"grpc-port"`
	GrpcHost          string   `mapstructure:"grpc-host"`
	HttpPort          int      `mapstructure:"http-port"`
	HttpHost          string   `mapstructure:"http-host"`
	AllowedOrigins    []string `mapstructure:"allowed-origins"`
	Web3Provider      string   `mapstructure:"web3-provider"`
	PrivateKey        string   `mapstructure:"private-key"`
	FundingAmount     string   `mapstructure:"funding-amount"`
	GasLimit          uint64   `mapstructure:"gas-limit"`
	IpLimitPerAddress int      `mapstructure:"ip-limit-per-address"`
	ChainId           int64    `mapstructure:"chain-id"`
}

// Server capable of funding requests for faucet ETH via gRPC and REST HTTP.
type Server struct {
	faucetpb.UnimplementedFaucetServer
	cfg           *Config
	client        *ethclient.Client
	funder        common.Address
	pk            *ecdsa.PrivateKey
	fundingAmount *big.Int
}

// NewServer initializes the server from configuration values.
func NewServer(cfg *Config) (*Server, error) {
	privKeyHex := cfg.PrivateKey
	if strings.HasPrefix(privKeyHex, "0x") {
		privKeyHex = privKeyHex[2:]
	}
	pk, err := crypto.HexToECDSA(privKeyHex)
	if err != nil {
		return nil, fmt.Errorf("could not parse funder private key: %v", err)
	}
	fundingAmount, ok := new(big.Int).SetString(cfg.FundingAmount, 10)
	if !ok {
		return nil, errors.New("could not set funding amount")
	}
	client, err := ethclient.DialContext(context.Background(), cfg.Web3Provider)
	if err != nil {
		return nil, fmt.Errorf("could not dial %s: %w", cfg.Web3Provider, err)
	}
	funder := crypto.PubkeyToAddress(pk.PublicKey)
	return &Server{
		cfg:           cfg,
		client:        client,
		funder:        funder,
		pk:            pk,
		fundingAmount: fundingAmount,
	}, nil
}

// Start a faucet server by serving a gRPC connection, an http JSON server
func (s *Server) Start() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	runtime.GOMAXPROCS(runtime.NumCPU())

	log.WithFields(logrus.Fields{
		"chainID": s.cfg.ChainId,
	}).Info("Initializing faucet server")

	// Query the funds left in the funder's account.
	s.queryFundsLeft(ctx)

	// Initialize and register gRPC handlers.
	grpcServer := s.initializeGRPCServer()

	grpcAddress := fmt.Sprintf("%s:%d", s.cfg.GrpcHost, s.cfg.GrpcPort)
	// Start a gRPC server.
	go func() {
		log.Infof("Starting gRPC server %s", grpcAddress)
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.GrpcPort))
		if err != nil {
			log.WithError(err).Fatalf("Could not listen on port %d", s.cfg.GrpcPort)
		}
		if err := grpcServer.Serve(lis); err != nil {
			log.WithError(err).Fatal("Stopped server")
		}
	}()

	// Start a gRPC Gateway to serve http JSON requests.
	gatewayAddress := fmt.Sprintf("%s:%d", s.cfg.HttpHost, s.cfg.HttpPort)
	gatewaySrv := gateway.New(ctx, &gateway.Config{
		GatewayAddress:      gatewayAddress,
		RemoteAddress:       grpcAddress,
		AllowedOrigins:      s.cfg.AllowedOrigins,
		EndpointsToRegister: []gateway.RegistrationFunc{faucetpb.RegisterFaucetHandlerFromEndpoint},
	})
	log.Infof("Starting JSON http server %s", gatewayAddress)
	gatewaySrv.Start()

	// Listen for any process interrupts.
	stop := make(chan struct{})
	go func() {
		sigc := make(chan os.Signal, 1)
		signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(sigc)
		<-sigc
		logrus.Info("Got interrupt, shutting down...")
		grpcServer.GracefulStop()
		stop <- struct{}{}
	}()

	// Wait for stop channel to be closed.
	<-stop
}

// Query the funds left in the faucet account and log them to the uer.
func (s *Server) queryFundsLeft(ctx context.Context) {
	bal, err := s.client.BalanceAt(ctx, s.funder, nil)
	if err != nil {
		log.WithError(err).Fatalf("Could not retrieve funder's current balance")
	}

	log.WithFields(logrus.Fields{
		"fundsInWei": bal,
		"publicKey":  s.funder.Hex(),
	}).Info("Funder account details")
}

// Initialize a gRPC server and register handlers.
func (s *Server) initializeGRPCServer() *grpc.Server {
	grpcServer := grpc.NewServer()
	faucetpb.RegisterFaucetServer(grpcServer, s)
	reflection.Register(grpcServer)
	return grpcServer
}
