/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/nambrosini/scribe/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var cfg config.AppConfig

var (
	modelType string
	url       string
	apiKey    string
	modelName string
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
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&modelType, "modelType", "m", "mistral", "the model it should use (mistral|ollama)")
	viper.BindPFlag("modelType", rootCmd.PersistentFlags().Lookup("model.modelType"))

	rootCmd.PersistentFlags().StringVarP(&url, "url", "u", "https://api.mistral.ai/v1/chat/completions", "the url to the models rest api")
	viper.BindPFlag("url", rootCmd.PersistentFlags().Lookup("model.url"))

	rootCmd.PersistentFlags().StringVarP(&apiKey, "key", "k", "", "the key to use to authenticate against the rest api")
	viper.BindPFlag("key", rootCmd.PersistentFlags().Lookup("model.apiKey"))
	viper.BindEnv("model.apiKey", "MODEL_API_KEY")

	rootCmd.PersistentFlags().StringVarP(&modelName, "name", "n", "mistral-large-latest", "the model that should be used")
	viper.BindPFlag("name", rootCmd.PersistentFlags().Lookup("model.name"))

}

func initConfig() {
	viper.SetConfigType("toml")
	viper.SetConfigName("scribe")
	viper.AddConfigPath("$XDG_CONFIG_HOME")

	// viper.SetEnvPrefix("SCRIBE")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
	} else {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	err := viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
}
