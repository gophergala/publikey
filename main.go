package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/gerred/publikey/publikey"
)

func main() {
	app := cli.NewApp()
	app.Name = "publikey"
	app.Author = Author
	app.Email = Email
	app.Version = Version
	app.Usage = "interact with the publikey API for public SSH key management"
	app.Commands = []cli.Command{
		publikey.NewListCommand(),
	}

	app.Run(os.Args)
}
