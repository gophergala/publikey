package publikey

import (
	"fmt"

	"github.com/codegangsta/cli"
)

func NewListCommand() cli.Command {
	return cli.Command{
		Name:   "ls",
		Usage:  "list keys for a user, defaults to logged in user",
		Action: handleListCommand,
	}
}

func handleListCommand(c *cli.Context) {
	fmt.Println("that's how you get ants, barry")
}
