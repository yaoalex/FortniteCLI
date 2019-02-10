package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd creates a new cobra command
var RootCmd = &cobra.Command{
	Use:   "FortniteCLI",
	Short: "FortniteCLI is a command line interface to track your Fortnite stats",
}
