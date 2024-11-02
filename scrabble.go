package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var people int64

func main() {
	app := &cli.App{
		Name: "scrabble",
		Flags: []cli.Flag{
			&cli.Int64Flag{
				Name:        "num_of_people",
				Value:       2,
				Aliases:     []string{"n"},
				Usage:       "number of people playing scrabble",
				Required:    true,
				Destination: &people,
			},
		},
		Action: game,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
