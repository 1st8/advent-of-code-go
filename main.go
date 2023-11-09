/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"

	"github.com/1st8/advent-of-code-go/cmd"
	"github.com/spf13/viper"
)

func main() {
	// Initialize Viper to read the configuration
	viper.SetConfigName("config")     // Name of config file (without extension)
	viper.SetConfigType("yaml")       // Required if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/.aoc") // Call multiple times to add many search paths
	viper.AddConfigPath(".")          // Optionally look for config in the working directory

	// Read in the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read config file: %v\n", err)
		os.Exit(1)
	}

	// Execute the command
	// if err := cmd.Execute(); err != nil {
	// 	fmt.Fprintf(os.Stderr, "Command execution failed: %v\n", err)
	// 	os.Exit(1)
	// }
	cmd.Execute()
}
