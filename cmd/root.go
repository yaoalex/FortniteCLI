package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd creates a new cobra command
var RootCmd = &cobra.Command{
	Use:   "fortnitecli",
	Short: "FortniteCLI is a command line interface for tracking your Fortnite stats",
}
