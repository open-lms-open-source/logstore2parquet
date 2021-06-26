package converter

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Command triggers the conversion process.
func Command() *cobra.Command {
	c := &cobra.Command{
		Use:   "convert",
		Short: "Convert a CSV",
		Long:  `Convert a CSV to Parquet format`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("requires an input and output file argument")
			}
			if _, err := os.Stat(args[0]); err != nil {
				return fmt.Errorf("cannot read input file")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			doConvert(args[0], args[1])
			return nil
		},
	}

	// Custom template to display our required args in the usage text.
	c.SetUsageTemplate(`Usage:{{if .Runnable}}
  {{.UseLine}} <inputfile> <outputfile>{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if .HasAvailableLocalFlags}}
Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}
Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}
  `)

	return c
}
