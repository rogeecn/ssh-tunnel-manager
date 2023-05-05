package cmd

import (
	"log"
	"os"

	"github.com/rogeecn/stm/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var conf config.Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stm",
	Short: "ssh tunnel manager",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		viper.SetConfigFile(cfgFile)

		if err := viper.ReadInConfig(); err != nil {
			log.Fatal(err)
		}

		if err := viper.Unmarshal(&conf); err != nil {
			log.Fatal(err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", ".stm.toml", "config file (default is $HOME/.stm.toml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
