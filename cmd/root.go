package cmd

import (
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "ritexfmt [text]",
    Short: "Minimal text formatter for fun.",
    Long:  `ritexfmt is a CLI tool for formatting and analyzing text. ...`,
    Args:  cobra.ArbitraryArgs,
    RunE:  run,
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}
