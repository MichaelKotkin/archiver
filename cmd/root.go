package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "Archiver",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		HandlerErr(err)
	}
}

func HandlerErr(err error) {
	fmt.Println(err)
	os.Exit(1)
}
