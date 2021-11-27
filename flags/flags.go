package flags

import "github.com/urfave/cli"

var Flags = make([]cli.Flag, 0)

func init() {
	Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "ip",
			Value: "127.0.0.1",
			Usage: "ip of nodered",
		},
		cli.IntFlag{
			Name:  "port, p",
			Value: 1880,
			Usage: "port of nodered",
			//EnvVar: "",
		},
		cli.StringFlag{
			Name:        "path",
			Usage:       "path of flows.json",
			Value:       "C:\\Users\\weibin\\.node-red",
			Destination: nil,
		},
	}
}
