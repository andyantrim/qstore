package main

import (
	"fmt"
	"log"
	"os"

	"github.com/andyantrim/qstore/shell"
	"github.com/andyantrim/qstore/store/cache"
	"github.com/andyantrim/qstore/transport/rest"
	"github.com/urfave/cli/v2"
)

func main() {
	store := cache.NewCache()
	app := cli.App{
		Name:      "qstore",
		UsageText: "A simple key-value store",
		Action: func(c *cli.Context) error {
			fmt.Println(c.App.Name)
			fmt.Println(c.App.UsageText)
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "interactive",
				Aliases: []string{"i"},
				Usage:   "Start the interactive mode",
				Action: func(c *cli.Context) error {
					session := shell.NewShell(store)
					session.Run()
					fmt.Println("Interactive mode session closed")
					return nil
				},
			},
			{
				Name:    "serve",
				Aliases: []string{"s"},
				Usage:   "Start the server mode",
				Action: func(c *cli.Context) error {
					server := rest.NewServer(store)
					fmt.Println("Server mode session started on port 16767")
					if err := server.ListenAndServe(":16767"); err != nil {
						log.Fatal(err)
					}
					fmt.Println("Server mode session closed")
					return nil
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
