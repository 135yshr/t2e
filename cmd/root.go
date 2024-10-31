/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

//nolint:forbidigo,revive // Output the results to the terminal
func newRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "t2e",
		Version: Version(),
		Short:   "This program converts the specified time to epoch time.",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(time.Now().Unix())

			return nil
		},
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := newRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
