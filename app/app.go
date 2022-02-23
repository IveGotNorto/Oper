package app

import (
	"fmt"
	"oper/store"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func Run(store store.PasswordStore) int {

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
			return store.TreeList()
		},
		Before: func(c *cli.Context) error {
			return store.Setup()
		},
		Commands: []*cli.Command{
			{
				Name:        "ls",
				Aliases:     []string{"list"},
				Description: "List passwords from the One Password command line utility",
				Action: func(c *cli.Context) error {
					return store.TreeList()
				},
			},
			{
				Name:        "upls",
				Aliases:     []string{"unpretty-list"},
				Description: "List passwords, with no formatting, from the One Password command line utility",
				Action: func(c *cli.Context) error {
					return store.List()
				},
			},
			{
				Name:        "show",
				Description: "Print the password under the password-name",
				Action: func(c *cli.Context) error {
					if c.Args().Len() >= 1 {
						store.Show(c.Args().First())
					}
					return nil
				},
				ArgsUsage: "password-name",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "copy"},
				},
			},
			{
				Name:        "find",
				Aliases:     []string{"search"},
				Description: "List names of passwords and vaults that match pass-names",
				Action: func(c *cli.Context) error {
					if c.Args().Present() {
						store.Find(c.Args().Slice())
					}
					return nil
				},
				ArgsUsage: "pass-names...",
			},
			{
				Name:        "insert",
				Aliases:     []string{"add"},
				Description: "",
				Action: func(c *cli.Context) error {
					if c.Args().Present() {
						return store.Insert(c.Args().First())
					}
					return nil
				},
				ArgsUsage: "pass-name",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "echo"},
					&cli.BoolFlag{Name: "multiline"},
					&cli.BoolFlag{Name: "force"},
				},
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "debug"},
		},
	}

	if err := cli.Run(os.Args); err != nil {
		fmt.Printf("%s\n", err)
		return 1
	}
	return 0
}
