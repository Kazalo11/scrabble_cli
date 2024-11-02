package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	var people int64
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
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
