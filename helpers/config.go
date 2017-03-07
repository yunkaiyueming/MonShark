package helpers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
)

func GetMyConfig(file, item string) (string, error) {
	need, err := isNeedDiffByMachine(file)
	fmt.Println(need)
	if err != nil {
		return "", err
	}

	if need {
		machineId, err := getMyMachineId()
		if err != nil {
			return "", err
		}
		fmt.Println("conf/" + machineId + "/" + file + ".conf")
		cnf, err := config.NewConfig("ini", "conf/"+machineId+"/"+file+".conf")
		if err != nil {
			return "", ErrLog("parse conf/machines/%s/%s.conf fail %s", machineId, file, err.Error())
		}

		return cnf.String(item), nil
	}

	return beego.AppConfig.String(item), nil
}

func isNeedDiffByMachine(file string) (bool, error) {
	cnf, err := config.NewConfig("ini", "conf/need_diff.conf")
	if err != nil {
		return false, ErrLog("parse conf/need_diff.conf fail %s", err.Error())
	}

	isNeed, err := cnf.Bool(file)
	if err != nil {
		return false, ErrLog("parse conf/need_diff.conf fail %s", err.Error())
	}

	if isNeed {
		return true, nil
	} else {
		return false, nil
	}
}

func getMyMachineId() (string, error) {
	cnf, err := config.NewConfig("ini", "E:/machine.conf")
	if err != nil {
		return "", ErrLog("parse /etc/rayjoy_plattech/machine.conf fail %s", err.Error())
	}

	return cnf.String("machine_id"), nil
}

func ErrLog(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}
