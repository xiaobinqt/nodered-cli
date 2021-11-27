package comm

import (
	"fmt"
	"os"

	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"nodered-cli/model"
)

func filepath(dirPath string) string {
	return fmt.Sprintf("%s/flows.json", dirPath)
}

func writeFile(dirPath string) (err error) {
	f, err := os.OpenFile(filepath(dirPath), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	defer f.Close()

	if err != nil {
		err = errors.Wrapf(err, "OpenFile err")
		logrus.Error(err.Error())
		return err
	}

	_, err = f.WriteString("[]")
	if err != nil {
		err = errors.Wrapf(err, "file WriteString err")
		logrus.Error(err.Error())
		return err
	}

	return nil
}

func ParseFlowsJSON(dirPath string) (jsonVal []model.FlowItem, err error) {
	jsonVal = make([]model.FlowItem, 0)

	filePtr, err := os.Open(filepath(dirPath))
	if err != nil {
		err = errors.Wrapf(err, "open file ParseFlowsJSON err")
		logrus.Error(err.Error())
		return jsonVal, err
	}
	defer filePtr.Close()

	decoder := jsoniter.NewDecoder(filePtr)
	err = decoder.Decode(&jsonVal)
	if err != nil {
		err = errors.Wrapf(err, "ParseFlowsJSON Decode err")
		logrus.Error(err.Error())
		return jsonVal, err
	}

	logrus.Debugf("ParseFlowsJSON success: %+v \n", jsonVal)
	return jsonVal, err
}

func FlowsJSON(dirPath string) (flowsVal []model.FlowItem, err error) {
	flowsVal = make([]model.FlowItem, 0)

	_, err = os.Stat(filepath(dirPath))
	// 文件不存在，创建文件并写入 []
	if err != nil && os.IsExist(err) == false {
		err = writeFile(filepath(dirPath))
		if err != nil {
			return flowsVal, err
		}
		return flowsVal, nil
	}

	// 文件存在读取文件
	flowsVal, err = ParseFlowsJSON(dirPath)
	if err != nil {
		return flowsVal, err
	}

	return flowsVal, nil
}
