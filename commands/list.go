package commands

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"nodered-cli/comm"
	"nodered-cli/flags"
	"nodered-cli/model"
)

var ListCmd = cli.Command{
	Name:   "list",
	Usage:  "nodered label list also is project list",
	Flags:  flags.Flags,
	Action: listAction,
}

func listAction(cli *cli.Context) {
	var (
		flowJsonVal = make([]model.FlowItem, 0)
		err         error
	)

	flowJsonVal, err = comm.ParseFlowsJSON(cli.String("path"))
	if err != nil {
		err = errors.Wrapf(err, "ParseFlowsJSON err")
		logrus.Fatal(err.Error())
	}

	for _, each := range flowJsonVal {
		if each.Type == "tab" {
			fmt.Printf("projectName: %s  \n", each.Label)
		}
	}
	//fmt.Println(flowJsonVal)
}
