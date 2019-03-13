package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/loqwai/crispy-modulus/game"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "crispy-modulus",
	Short: "crispy-modulus is a card game",
	RunE: func(cmd *cobra.Command, args []string) error {
		g := game.New(3)

		for {
			err := printState(g.GetState())
			if err != nil {
				return err
			}

			command, err := getCommand()
			if err != nil {
				return err
			}

			if command == "d" {
				err = g.Draw()
				if err != nil {
					return err
				}
				continue
			}

			if strings.HasPrefix(command, "s") {
				val, err := strconv.Atoi(strings.TrimPrefix(command, "s"))
				if err != nil {
					return err
				}

				err = g.Steal(val)
				if err != nil {
					return err
				}
			}
		}
	},
}

func getCommand() (string, error) {
	_, err := fmt.Println("Whatcha wanna do? (d: draw, s#: steal card)")
	if err != nil {
		return "", err
	}

	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(command), nil
}

func printState(state game.State) error {
	output, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}

	_, err = fmt.Println(string(output))
	return err
}

// Execute is the main runtime of the application
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
