package cmd

import (
	"cargorun/core"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var exampleCmd = &cobra.Command{
	Use:   "example",
	Short: "run a cargo example",
	Args:  cobra.MinimumNArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		examples, err := core.GetRunnable("example")
		if err != nil {
			log.Fatal(err)
		}
		return examples, cobra.ShellCompDirectiveNoFileComp
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		examples, err := core.GetRunnable("example")
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return nil
		}

		if len(examples) == 0 {
			fmt.Fprintf(os.Stderr, "no examples found\n")
			return nil
		}

		example := args[0]
		extraArgs := args[1:]
		validExample := false
		for _, e := range examples {
			if example == e {
				validExample = true
				break
			}
		}
		if !validExample {
			fmt.Fprintf(os.Stderr, "invalid example %v\n", example)
			return nil
		}

		err = core.Run("example", example, extraArgs)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return nil
		}

		return nil

	},
}

func init() {
	rootCmd.AddCommand(exampleCmd)
}
