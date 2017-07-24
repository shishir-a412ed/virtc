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
			Usage:       "Pull an OCI compliant image from the registry",
			Description: "Pull an OCI compliant image from the registry",
			Action: func(context *cli.Context) error {
				if err := pullImage(context); err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:        "push",
			Usage:       "Push an OCI compliant image to the registry",
			Description: "Push an OCI compliant image to the registry",
		},
	},
}

func pullImage(context *cli.Context) error {
	fmt.Println("Pulling image")
	return nil
}
