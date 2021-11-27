package comm

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestFlowsJSON(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	ret, err := FlowsJSON("C:\\Users\\weibin\\.node-red")
	if err != nil {
		t.Errorf("err = %s ", err.Error())
		return
	}

	t.Logf("%+v", ret)
}
