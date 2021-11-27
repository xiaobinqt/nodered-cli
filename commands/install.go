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

var Install = cli.Command{
	Name:   "install, i",
	Usage:  "install project from tar package, tar package include project necessary to run node_modules",
	Flags:  flags.Flags,
	Action: InstallAction,
}

func InstallAction(cli *cli.Context) {
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
