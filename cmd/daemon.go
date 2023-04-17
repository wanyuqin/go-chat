package cmd

import (
	"github.com/spf13/cobra"

	"go-chat/initialize"
)

var daemonCommand = &cobra.Command{
	Use: "daemon",
	Run: daemonStart,
}

func init() {
	rootCmd.AddCommand(daemonCommand)
}

func daemonStart(cmd *cobra.Command, args []string) {
	initialize.InitRouter()

}
