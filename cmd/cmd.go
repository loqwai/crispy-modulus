package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/loqwai/crispy-modulus/game"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "crispy-modulus",
	Short: "crispy-modulus is a card game",
	RunE: func(cmd *cobra.Command, args []string) error {
		g := game.New()
		s := g.GetState()
		output, err := json.MarshalIndent(s, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

// Execute is the main runtime of the application
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
