package cmd

import (
	"fmt"
	"os"

	"github.com/loqwai/crispy-modulus/game"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "crispy-modulus",
	Short: "crispy-modulus is a card game",
	Run: func(cmd *cobra.Command, args []string) {
		g := game.New()
		fmt.Println(g.String())
	},
}

// Execute is the main runtime of the application
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
