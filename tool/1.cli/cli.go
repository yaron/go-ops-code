package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "url, u",
			Value: "https://github.com/yaron/go-ops.git",
			Usage: "URL to git clone",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "clone",
			Aliases: []string{"c"},
			Usage:   "Clone a git url",
			Action: func(c *cli.Context) error {
				fmt.Printf("URL: %q", c.String("url"))
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
