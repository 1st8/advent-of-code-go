/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	cli "github.com/1st8/advent-of-code-go/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	Use:   "init [day]",
	Short: "Initialize Advent of Code day",
	Long:  `This command initializes the input and setup for a given day for Advent of Code.`,
	Args:  cobra.ExactArgs(1),
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	day, err := strconv.Atoi(args[0])
	selectedYear, err := cmd.Flags().GetInt("year")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid day number: %v\n", err)
		os.Exit(1)
	}

	// Assuming your session is stored under the key "session" in a config file.
	session := viper.GetString("session")
	if session == "" {
		fmt.Fprintln(os.Stderr, "Session not found in configuration")
		os.Exit(1)
	}

	err = cli.DownloadInput(selectedYear, day, session)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error downloading input: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
}
