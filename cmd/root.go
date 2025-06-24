package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kk",
	Short: "Encrypted key/value store using age",
}

func Execute() {
	// Args 1 is the binary itself
	if len(os.Args) == 1 {
		// os.Args[0] bleibt der Binary-Name.
		// os.Args[1] wird der Name des Standardbefehls ("pick").
		rootCmd.SetArgs([]string{"pick"})
	}

	// Falls der Hilfe-Flag (--help oder -h) als einziges Argument übergeben wird,
	// sollte die Standardaktion NICHT ausgelöst werden, sondern die Hilfe angezeigt werden.
	// Dieser Teil ist etwas tricky, da Cobra die Argumente selbst parst.
	// Eine robustere Lösung ist, Cobra die Argumente parsen zu lassen
	// und dann zu prüfen, ob ein Befehl gefunden wurde.
	cmd, _, err := rootCmd.Find(os.Args[1:])

	// Wenn kein Befehl gefunden wurde ODER der gefundene Befehl der Root-Befehl selbst ist
	// (was der Fall ist, wenn keine Sub-Kommandos angegeben wurden), UND es nicht der Hilfe-Befehl ist.
	if err != nil || cmd.Use == rootCmd.Use {
		// Überprüfe, ob --help oder -h vorhanden ist
		helpFlagGiven := false
		for _, arg := range os.Args {
			if arg == "--help" || arg == "-h" {
				helpFlagGiven = true
				break
			}
		}

		// Setze das Standardkommando nur, wenn kein Befehl angegeben wurde
		// UND KEIN Hilfe-Flag angegeben wurde.
		if !helpFlagGiven && len(os.Args) > 1 && !isSubcommandPresent(rootCmd, os.Args[1:]) {
			// Setze die Argumente so, dass sie den "pick"-Befehl aufrufen
			// und die restlichen Argumente als Argumente für "pick" weiterleiten.
			newArgs := append([]string{"pick"}, os.Args[1:]...)
			rootCmd.SetArgs(newArgs)
		}
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// checks if any of the provided args is a known subcommand of the root command.
func isSubcommandPresent(rootCmd *cobra.Command, args []string) bool {
	for _, arg := range args {
		for _, sub := range rootCmd.Commands() {
			if sub.Name() == arg || sub.HasAlias(arg) {
				return true
			}
		}
	}
	return false
}
