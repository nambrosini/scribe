/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "scribe",
	Short: "scribe is a tool that automates commit message generation by sending git diffs to an LLM, ensuring concise and meaningful messages for your version control.",
	Long:  `scribe is an innovative tool designed to streamline your version control process by automatically generating commit messages. By sending your git diff output to a Large Language Model (LLM), scribe crafts concise and meaningful commit messages that accurately reflect your changes. This not only saves you time but also helps maintain a clear and consistent commit history, enhancing collaboration and project management. Perfect for developers looking to optimize their workflow with the power of AI, scribe integrates seamlessly into your existing development environment.`,
	// Run: func(cmd *cobra.Command, args []string) {
	//
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.scribe.yaml)")
}
