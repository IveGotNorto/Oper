package main

import (
	"log"
	items "oper/items"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	var items items.Items
	//var vaults vaults.Vaults

	app := &cli.App{
		Name:     "Oper",
		Version:  "v0.0.1",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Norto",
				Email: "ivegot@norto.dev",
			},
		},
		Usage: "One Password command line wrapper",
		Commands: []*cli.Command{
			{
				Name:        "ls",
				Aliases:     []string{"list"},
				Description: "List passwords in the One Password Command line utility",
				Action: func(c *cli.Context) error {
					OpPrint(&items)
					return nil
				},
				Before: func(c *cli.Context) error {
					return items.Retrieve()
				},
			},
			{
				Name:        "show",
				Description: "Print the password under the password-name",
				Action: func(c *cli.Context) error {
					if c.Args().Len() >= 1 {
						return OpShow(&items, c.Args().First())
					}
					return nil
				},
				ArgsUsage: "password-name",
				Before: func(c *cli.Context) error {
					return items.Retrieve()
				},
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
