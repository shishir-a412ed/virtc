package main

import (
	"fmt"
	"github.com/nvsandbox/cmd"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

func main() {
	var logLevel, logFormat string

	app := cli.NewApp()
	app.Name = "nvidia-containerizer"
	app.Usage = "Containerize deep learning (DL) jobs."
	app.Version = "0.0.1"
	app.EnableBashCompletion = true
	app.Authors = []cli.Author{
		{
			Name:  "Shishir Mahajan",
			Email: "shishirm@nvidia.com",
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "log-level",
			Value:       "warn",
			Usage:       "Set the log level: debug, info, warn, error, fatal, panic",
			Destination: &logLevel,
		},
		cli.StringFlag{
			Name:        "log-format",
			Value:       "text",
			Usage:       "Set the log format: text or json",
			Destination: &logFormat,
		},
	}

	app.Commands = []cli.Command{
		cmd.ImageCommand,
		cmd.ContainerCommand,
		cmd.VolumeCommand,
	}

	app.Before = func(c *cli.Context) error {
		// Set the log format.
		switch logFormat {
		case "text":
			// Do nothing. Since {text} is default logrus format.
		case "json":
			log.SetFormatter(&log.JSONFormatter{})
		default:
			return fmt.Errorf("Unknown log-format %q\n", logFormat)
		}

		// Set the log level.
		switch logLevel {
		case "debug":
			log.SetLevel(log.DebugLevel)
		case "info":
			log.SetLevel(log.InfoLevel)
		case "warn":
			log.SetLevel(log.WarnLevel)
		case "error":
			log.SetLevel(log.ErrorLevel)
		case "fatal":
			log.SetLevel(log.FatalLevel)
		case "panic":
			log.SetLevel(log.PanicLevel)
		default:
			return fmt.Errorf("Unknown log-level %q\n", logLevel)
		}
		return nil
	}

	app.Action = func(c *cli.Context) error {
		fmt.Printf("Log level is %s\n", logLevel)
		fmt.Printf("Log format is %s\n", logFormat)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
