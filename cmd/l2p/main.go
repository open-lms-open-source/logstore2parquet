package main

import (
	"github.com/spf13/cobra"

	"logstore2parquet/internal/converter"
	"logstore2parquet/internal/version"
)

func main() {
	//version.App = "l2p"

	var command = &cobra.Command{
		Use:   "l2p",
		Short: "Generage parquet files from Moodle™ logstore csv exports",
	}
	command.AddCommand(converter.Command())
	command.AddCommand(version.Command())

	command.Execute()
}
