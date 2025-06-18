package cmd

import (
	"fmt"

	"github.com/moosemanf/kk/vault"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get <key>",
	Short: "Get a value from the vault",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		v, err := vault.LoadVault()
		if err != nil {
			return err
		}
		val, ok := v[args[0]]
		if !ok {
			fmt.Println("Key not found.")
			return nil
		}
		fmt.Println(val)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
