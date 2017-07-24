package cmd

import (
	"fmt"
	"github.com/urfave/cli"
)

var ImageCommand = cli.Command{
	Name:        "images",
	Usage:       "OCI compliant image management",
	Description: "images sub-commands manages OCI compliant images",
	Subcommands: []cli.Command{
		{
			Name:        "pull",
			Usage:       "nvsandbox images pull [OPTIONS] IMAGE",
			Description: "Pull an OCI compliant image from the registry",
		},
		{
			Name:        "push",
			Usage:       "nvsandbox images push [OPTIONS] IMAGE",
			Description: "Push an OCI compliant image to the registry",
		},
	},

	Action: func(context *cli.Context) error {
		fmt.Println("HELLO IMAGE COMMAND")
		return nil
	},
}
