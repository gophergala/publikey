package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "publikey"
	app.Author = Author
	app.Email = Email
	app.Version = Version
	app.Usage = "interact with the publikey API for public SSH key management"

	app.Run(os.Args)
}
