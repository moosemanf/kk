package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/moosemanf/kk/vault"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set <key>",
	Short: "Set a key in the vault",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]

		fmt.Print("Value: ")
		reader := bufio.NewReader(os.Stdin)
		value, _ := reader.ReadString('\n')
		value = strings.TrimSpace(value)

		v, err := vault.LoadVault()
		if err != nil {
			return err
		}
		v[key] = value
		return vault.SaveVault(v)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
