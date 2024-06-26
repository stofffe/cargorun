package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/stofffe/cargorun/core"

	"github.com/spf13/cobra"
)

var binCmd = &cobra.Command{
	Use:   "bin",
	Short: "run a cargo binary",
	Args:  cobra.MinimumNArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		binaries, err := core.GetRunnable("bin")
		if err != nil {
			log.Fatal(err)
		}
		return binaries, cobra.ShellCompDirectiveNoFileComp
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		binaries, err := core.GetRunnable("bin")
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return nil
		}

		if len(binaries) == 0 {
			fmt.Fprintf(os.Stderr, "no binaries found\n")
			return nil
		}

		bin := args[0]
		extraArgs := args[1:]
		validBin := false
		for _, e := range binaries {
			if bin == e {
				validBin = true
				break
			}
		}
		if !validBin {
			fmt.Fprintf(os.Stderr, "invalid binary %v\n", bin)
			return nil
		}

		err = core.Run("bin", bin, extraArgs)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return nil
		}

		return nil

	},
}

func init() {
	rootCmd.AddCommand(binCmd)
}
