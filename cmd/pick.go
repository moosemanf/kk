package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/atotto/clipboard"
	"github.com/moosemanf/kk/vault"
	"github.com/spf13/cobra"
)

var clip bool

var pickCmd = &cobra.Command{
	Use:   "pick",
	Short: "Fuzzy-pick a key using fzf and get its value",
	RunE: func(cmd *cobra.Command, args []string) error {
		v, err := vault.LoadVault()
		if err != nil {
			return err
		}
		if len(v) == 0 {
			fmt.Println("Vault is empty.")
			return nil
		}

		// Create fzf input
		var buf bytes.Buffer
		for k := range v {
			buf.WriteString(k + "\n")
		}

		// Run fzf
		fzf := exec.Command("fzf", "--prompt=🔑 Pick key: ")
		fzf.Stdin = &buf
		fzf.Stderr = os.Stderr
		out, err := fzf.Output()
		if err != nil {
			// A non-zero exit code from fzf usually means the user aborted
			// by pressing Ctrl-C or Esc. We can safely return nil here.
			if _, ok := err.(*exec.ExitError); ok {
				return nil
			}
			// For other errors, return them.
			return fmt.Errorf("fzf execution failed: %w", err)
		}

		key := string(bytes.TrimSpace(out))
		// If the key is empty, the user likely aborted without making a selection.
		if key == "" {
			return nil
		}

		val, ok := v[key]
		if !ok {
			return fmt.Errorf("selected key '%s' not found in vault", key)
		}

		if clip {
			if err := clipboard.WriteAll(val); err != nil {
				return err
			}
			fmt.Printf("📋 Copied value of '%s' to clipboard.\n", key)
		} else {
			fmt.Println(val)
		}

		return nil
	},
}

func init() {
	pickCmd.Flags().BoolVarP(&clip, "clip", "c", false, "Copy value to clipboard")
	rootCmd.AddCommand(pickCmd)
}
