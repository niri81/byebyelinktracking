package cmd

import (
	"os"

	"github.com/niri81/byebyelinktracking/pkg/byebyelinktracking"
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "byebyelinktracking",
	Short: "Remove link tracking from links in your clipboard",
	Long:  `"Bye Bye, Link Tracking" removes link tracking data from links that are in your clipboard`,
	Run: func(cmd *cobra.Command, args []string) {
		byebyelinktracking.Run(cfgFile)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "p", "./config.json", "config file")
}
