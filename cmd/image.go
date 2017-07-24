package cmd

import (
	"fmt"
	"github.com/urfave/cli"
)

var ImageCommand = cli.Command{
	Name:        "images",
	Usage:       "OCI compliant image management",
	Description: "images sub-commands manages OCI compliant images",

	Action: func(context *cli.Context) error {
		fmt.Println("HELLO IMAGE COMMAND")
		return nil
	},
}
