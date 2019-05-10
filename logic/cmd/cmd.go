package cmd

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/loqwai/crispy-modulus/game"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "crispy-modulus",
	Short: "crispy-modulus is a card game",
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		rand.Seed(time.Now().UTC().UnixNano())
		cardCount := 5
		if len(args) > 0 {
			cardCount, err = strconv.Atoi(args[0])
			if err != nil {
				return err
			}
		}
		g := game.New(cardCount)
		g.Start()
		for {
			if g.IsDone() {
				break
			}

			err := printGame(g)
			if err != nil {
				return err
			}

			command, err := getCommand()
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				continue
			}

			if command == "d" {
				err = g.Draw()
				if err != nil {
					fmt.Fprintln(os.Stderr, err.Error())
					continue
				}
				continue
			}

			if strings.HasPrefix(command, "s") {
				val, err := strconv.Atoi(strings.TrimPrefix(command, "s"))
				if err != nil {
					fmt.Fprintln(os.Stderr, err.Error())
					continue
				}

				err = g.Steal(val)
				if err != nil {
					fmt.Fprintln(os.Stderr, err.Error())
					continue
				}
				continue
			}

			fmt.Fprintln(os.Stderr, "Unrecognized command: ", command)
		}

		fmt.Println("We have a winner! It was:", g.WhoIsWinning())
		return nil
	},
}

func getCommand() (string, error) {
	_, err := fmt.Fprintln(os.Stderr, "Whatcha wanna do? (d: draw, s#: steal card)")
	if err != nil {
		return "", err
	}

	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(command), nil
}

func printGame(g game.Game) error {
	output, err := g.String()
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
