package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "gchat",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
		return
	}
}
