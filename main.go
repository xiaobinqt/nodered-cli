package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	globalFlags = []cli.Flag{
		cli.StringFlag{
			Name:   "api-key",
			Usage:  "cSphere API key",
			EnvVar: "CSPHERE_API_KEY",
		},
		cli.StringFlag{
			Name:   "endpoint",
			Usage:  "cSphere API endpoint",
			EnvVar: "CSPHERE_API_ENDPOINT",
		},
	}

	commonFlags = []cli.Flag{
		cli.StringFlag{
			Name:  "app",
			Usage: "the name of the csphere APP",
		},
		cli.StringFlag{
			Name:  "service",
			Usage: "the name of the service",
		},
	}

	commonInstanceFlags = []cli.Flag{
		cli.StringFlag{
			Name:  "name",
			Usage: "instance name",
		},
	}

	cmdUpdateImageTag = cli.Command{
		Name:  "update-image-tag",
		Usage: "Update the service variables of a service",
		Flags: append(commonFlags,
			cli.StringFlag{
				Name:  "image-tag",
				Usage: "the new image tag to use",
			}),
		//Action: actionUpdateImageTag,
	}

	cmdSyncFiles = cli.Command{
		Name:  "sync-files",
		Usage: "Sync files to containers",
		Flags: append(commonFlags,
			cli.StringSliceFlag{
				Name:  "file, f",
				Usage: "files to sync. All changed files in the workspace by default",
			},
			cli.StringFlag{
				Name:  "dest",
				Value: "/",
				Usage: "the destination of the file in the container. Must end with '/' if it's a directory",
			},
			cli.StringFlag{
				Name:  "after-sync",
				Usage: "script to run in the container after files are synced",
			},
			cli.BoolFlag{
				Name:  "restart",
				Usage: "whether to restart the service after files are synced to the containers",
			}),
		//Action: actionSyncFiles,
	}

	cmdExec = cli.Command{
		Name:  "exec",
		Usage: "exec command in all the containers of the service",
		Flags: append(commonFlags,
			cli.StringFlag{
				Name:  "script",
				Usage: "the script to run",
			}),
		//Action: actionExec,
	}

	cmdCreateInstance = cli.Command{
		Name:  "create-instance",
		Usage: "create instance",
		Flags: append(commonInstanceFlags,
			cli.StringFlag{
				Name:  "pool-id",
				Usage: "svrpool id",
			},
			cli.StringFlag{
				Name:  "revision-id",
				Usage: "revision id",
			}),
		//Action: actionCreateInstance,
	}

	cmdUpgradeInstance = cli.Command{
		Name:  "upgrade-instance",
		Usage: "upgrade instance",
		Flags: append(commonInstanceFlags,
			cli.StringFlag{
				Name:  "new-revision-id",
				Usage: "new revision id",
			}),
		//Action: actionUpgradeInstance,
	}

	cmdGetInstance = cli.Command{
		Name:  "get-instance",
		Usage: "get instance info",
		Flags: commonInstanceFlags,
		//Action: actionGetInstance,
	}

	cmdDeleteInstance = cli.Command{
		Name:  "delete-instance",
		Usage: "delete instance",
		Flags: commonInstanceFlags,
		//Action: actionDeleteInstance,
	}

	cmdExpansionInstance = cli.Command{
		Name:  "extend-instance",
		Usage: "extend instance service",
		Flags: append(commonInstanceFlags,
			cli.StringFlag{
				Name:  "service-name",
				Usage: "service-name",
			}),
		//Action: actionExtendInstance,
	}

	cmdShrinkInstance = cli.Command{
		Name:  "shrink-instance",
		Usage: "shrink instance service",
		Flags: append(commonInstanceFlags,
			cli.StringFlag{
				Name:  "service-name",
				Usage: "service name",
			}),
		//Action: actionShrinkInstance,
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "nodered-cli"
	app.Usage = "nodered CLI tools"
	app.Version = "1.0.0"
	app.Author = "weibin"
	app.Email = "xiaobinqt@163.com"
	app.Flags = globalFlags

	app.Commands = []cli.Command{cmdUpdateImageTag, cmdSyncFiles, cmdExec, cmdCreateInstance, cmdUpgradeInstance, cmdGetInstance, cmdDeleteInstance, cmdExpansionInstance, cmdShrinkInstance}
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
