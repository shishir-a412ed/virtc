package cmd

import (
	"fmt"
	"github.com/urfave/cli"
)

var ContainerCommand = cli.Command{
	Name:        "containers",
	Usage:       "OCI compliant container management",
	Description: "containers sub-commands manages OCI compliant containers",

	Action: func(context *cli.Context) error {
		fmt.Println("HELLO CONTAINERS COMMAND")
		return nil
	},
}
