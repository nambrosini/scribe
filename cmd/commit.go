/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/nambrosini/scribe/internal/logic"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	issue        string
	mode         string
	templateFile string
	commitType   string
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit changes in repo with generated text",
	Long: `Commit takes the info passed via flags, sends the request with the requested template to the llm and starts the commit with the response.
	Note: before it commits, there is the possibility of making changes to the message.`,
	Run: func(cmd *cobra.Command, args []string) {
		msg, err := logic.SendRequest(cfg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = logic.Commit(msg)
		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)

	commitCmd.LocalFlags().StringVarP(&issue, "issue", "i", "", "the issue that should be incluted in the commit description (default is read from the ISSUE env variable)")
	viper.BindPFlag("commit.issue", commitCmd.Flags().Lookup("issue"))
	viper.BindEnv("commit.issue", "COMMIT_ISSUE")

	rootCmd.LocalFlags().StringVarP(&mode, "prompt", "p", "concise", "if the commit text will be concise (one line) or full (one line with text), default is concise")
	viper.BindPFlag("commit.prompt", commitCmd.Flags().Lookup("prompt"))

	rootCmd.LocalFlags().StringVarP(&templateFile, "templateFile", "f", "", "the file with the template to be used by the llm, the content will be sent directly to the llm (mode will be ignored if this flag is set)")
	viper.BindPFlag("commit.promptFile", commitCmd.Flags().Lookup("promptFile"))
}
