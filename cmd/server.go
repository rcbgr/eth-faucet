package cmd

import (
	"os"
	"runtime"
	"encoding/json"
	"github.com/rauljordan/eth-faucet/internal"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strconv"
)

type PrivateKey struct {
	Value string `json:"PrivateKey"`
}


var (
	log         = logrus.WithField("prefix", "cmd")
	cfgFilePath string
	rootCmd     = &cobra.Command{
		Use:   "faucet",
		Short: "Run a faucet server for Ethereum using captcha",
		RunE: func(command *cobra.Command, args []string) error {
			runtime.GOMAXPROCS(runtime.NumCPU())
			customFormatter := new(logrus.TextFormatter)
			customFormatter.TimestampFormat = "2006-01-02 15:04:05"
			logrus.SetFormatter(customFormatter)
			customFormatter.FullTimestamp = true

			var cfg *internal.Config
			if err := viper.Unmarshal(&cfg); err != nil {
				log.Fatal(err)
			}

			port, err := strconv.Atoi(os.Getenv("PORT"))
			if err != nil {
				log.Fatal("port not valid")
			}

			cfg.HttpPort = port

			cfg.Web3Provider = os.Getenv("WEB3_PROVIDER")

			if cfg.Web3Provider == "" {
				log.Fatal("WEB3_PROVIDER not set")
			}

			privateKeyJson := os.Getenv("PRIVATE_KEY")

			var privateKey PrivateKey
  		if err := json.Unmarshal([]byte(privateKeyJson), &privateKey); err != nil {
				log.Fatalf("PRIVATE_KEY problem: %v", err)
  		}

			cfg.PrivateKey = privateKey.Value

			if cfg.PrivateKey == "" {
				log.Fatal("--private-key hex string required")
			}
			srv, err := internal.NewServer(cfg)
			if err != nil {
				log.WithError(err).Fatal("Could not initialize faucet server")
			}
			srv.Start()
			return nil
		},
	}
)

func init() {

	cobra.OnInitialize(initConfig)

	rootCmd.Flags().StringVar(&cfgFilePath, "config", "", "Flag config yaml file path (optional)")
	rootCmd.Flags().Int("grpc-port", 5000, "Port to serve gRPC requests")
	rootCmd.Flags().String("grpc-host", "127.0.0.1", "Host to serve gRPC requests")
	//rootCmd.Flags().Int("http-port", 8000, "Port to serve REST http requests")
	rootCmd.Flags().String("http-host", "127.0.0.1", "Host to serve REST http requests")
	rootCmd.Flags().StringSlice("allowed-origins", []string{"*"}, "Allowed origins for REST http requests, comma-separated")
	//rootCmd.Flags().String("private-key", "", "Private key hex string of the funder of the faucet")
	rootCmd.Flags().String("funding-amount", "50000000000000000", "Amount in wei to fund with each request")
	rootCmd.Flags().Uint64("gas-limit", 29000000, "Gas limit for funding transactions")
	rootCmd.Flags().Int64("chain-id", 4, "Chain ID for Ethereum (4 is the rinkeby test network)")

	// Bind all flags to a viper configuration.
	if err := viper.BindPFlags(rootCmd.Flags()); err != nil {
		log.Fatal(err)
	}
	viper.SetDefault("author", "Raul Jordan <raul@prysmaticlabs.com>")
	viper.SetDefault("license", "MIT")
}

func initConfig() {
	// Use config file from the flag.
	viper.SetConfigFile(cfgFilePath)
	viper.AutomaticEnv()
	if cfgFilePath != "" {
		if err := viper.ReadInConfig(); err != nil {
			log.WithError(err).Fatalf("Could not read config file: %s", viper.ConfigFileUsed())
		}
	}
}

// Execute the faucet root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.WithError(err).Fatal("Could not execute root command")
		os.Exit(1)
	}
}
