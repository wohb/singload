package cmd

import (
	"fmt"
	"os"

	"github.com/wohb/singload/pkg/lb"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var addr string

var rootCmd = &cobra.Command{
	Use:   "singload",
	Short: "A load balancer without all of the complexity of handling multiple nodes",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize & Run the Singload load balancer
		singload := lb.LoadBalancer{
			ListenerPort: 80,
			TargetAddr:   addr,
		}
		singload.Run()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&addr, "address", "127.0.0.1", "the address of the node to which traffic will be routed")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.singload.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".singload" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".singload")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
