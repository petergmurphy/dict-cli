package cmd

import (
	"fmt"
	"os"

	"github.com/petergmurphy/cli-dict/pkg/dictionary"
	"github.com/petergmurphy/cli-dict/pkg/formatter"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "define",
  Short: "A dictionary CLI tool",
  Run: run,
}

func run(cmd *cobra.Command, args []string) {
    dict := dictionary.NewDictionary()
	definition, err := dict.Define(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	formatter := formatter.NewDictionaryFormatter()
	result := formatter.AddTitle(definition[0].Word).
	AddDefinitions(dictionary.GetWordDefintions(definition)).
	Build()

	fmt.Println(result)	
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
