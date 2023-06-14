package main

import (
	"fmt"
	"os"

	"github.com/154pinkchairs/loshell/cmd"
	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.App{
		Name: "loshell",
		Usage: `A CLI utility that allows to use some of the iterators
		from the samber/lo library on files.`,
		Commands: []*cli.Command{
			{
				Name:    "count",
				Aliases: []string{"c"},
				Usage:   "count operation",
				Action:  cmd.Count,
			},
			{
				Name:    "countby",
				Aliases: []string{"cb"},
				Usage:   "count by predicate (contains or empty, usage: loshell cb <file><predicate><string>)",
				Action:  cmd.CountBy,
			},
			{
				Name:    "subset",
				Aliases: []string{"s"},
				Usage:   "subset operation",
				Action:  cmd.Subset,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
