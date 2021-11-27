package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"nodered-cli/commands"
	"nodered-cli/flags"
)

var (
	err error
)

func main() {
	app := cli.NewApp()
	app.Name = "nodered-cli"
	app.Usage = "nodered CLI tools"
	app.Version = "1.0.0"
	app.Author = "weibin"
	app.Email = "xiaobinqt@163.com"
	app.Flags = flags.Flags
	app.Before = func(c *cli.Context) error {
		if c.String("path") == "" {
			err = fmt.Errorf("path is empty")
			logrus.Fatalf(err.Error())
		}

		if c.String("ip") == "" {
			logrus.Fatalf("ip empty")
		}

		if c.Int("port") <= 0 {
			logrus.Fatalf("port %d invalid", c.Int("port"))
		}

		return nil
	}

	app.Commands = []cli.Command{
		commands.ListCmd,
	}
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
