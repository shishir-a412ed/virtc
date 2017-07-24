package cmd

import (
	"fmt"
	"github.com/urfave/cli"
)

var VolumeCommand = cli.Command{
	Name:        "volumes",
	Usage:       "Volume management",
	Description: "volumes sub-commands manages ceph and nfs volumes",

	Action: func(context *cli.Context) error {
		fmt.Println("HELLO VOLUMES COMMAND")
		return nil
	},
}
