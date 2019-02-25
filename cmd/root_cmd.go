package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var configFile = ""

func init() {
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.env)")
}

// Execute executes commands to create the server
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "api",
	Short: "API",
	Long:  `REST Service`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("use: api serve")
	},
}
