package app

import (
	"fmt"
	"oper/store"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func Run(pass store.PasswordStore) int {

	cli := &cli.App{
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
		Action: func(c *cli.Context) error {
			return pass.TreeList("ascending")
		},
		Before: func(c *cli.Context) error {
			return pass.Setup(store.StoreArguments{
				Cache:   c.Bool("cache"),
				Verbose: c.Bool("verbose"),
				Debug:   c.Bool("debug"),
			})
		},
		Commands: []*cli.Command{
			{
				Name:        "ls",
				Aliases:     []string{"list"},
				Description: "List passwords from the One Password command line utility",
				Action: func(c *cli.Context) error {
					return pass.TreeList(c.String("order"))
				},
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "order", Aliases: []string{"o"}, Value: "ascending"},
				},
			},
			{
				Name:        "upls",
				Aliases:     []string{"unpretty-list"},
				Description: "List passwords, with no formatting, from the One Password command line utility",
				Action: func(c *cli.Context) error {
					return pass.List(c.String("order"))
				},
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "order", Aliases: []string{"o"}, Value: "ascending"},
				},
			},
			{
				Name:        "show",
				Description: "Print the password under the password-name",
				Action: func(c *cli.Context) error {
					if c.Args().Len() >= 1 {
						pass.Show(c.Args().First())
					}
					return nil
				},
				ArgsUsage: "password-name",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "copy"},
				},
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "debug"},
			&cli.BoolFlag{Name: "verbose"},
			&cli.BoolFlag{Name: "cache"},
		},
	}

	if err := cli.Run(os.Args); err != nil {
		fmt.Printf("%s\n", err)
		return 1
	}
	return 0
}
