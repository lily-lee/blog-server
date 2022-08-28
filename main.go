package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lily-lee/blog-server/cmd"
	"github.com/lily-lee/blog-server/config"
	"github.com/urfave/cli/v2"
)

func init() {
	config.Init()
}

// @title       blog api server
// @version     1.0
// @description This is a blog api server, serves api for blog web frontend.

// @contact.name  lily-lee
// @contact.email lilylee99.01@gmail.com

// @host     localhost:3000
// @BasePath /
func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "blog",
				Usage: "start blog server",
				Action: func(c *cli.Context) error {
					fmt.Println("start blog server ...")
					return cmd.BlogServer()
				},
			},
			{
				Name:  "migrate",
				Usage: "db migrate",
				Action: func(c *cli.Context) error {
					fmt.Println("start db migrate...")
					return cmd.Migrate()
				},
			},
			{
				Name:  "rollback",
				Usage: "db rollback",
				Action: func(c *cli.Context) error {
					fmt.Println("start db rollback...")
					return cmd.Rollback()
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
