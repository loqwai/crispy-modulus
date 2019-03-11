package cmd

import (
	"bufio"
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
		g := game.New(3)

		for {
			s := g.GetState()
			output, err := json.MarshalIndent(s, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			fmt.Println("Whatcha wanna do? (d: draw)")

			reader := bufio.NewReader(os.Stdin)
			_, err = reader.ReadString('\n')
			if err != nil {
				return err
			}
			err = g.Draw()
			if err != nil {
				return err
			}
		}
	},
}

// Execute is the main runtime of the application
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
