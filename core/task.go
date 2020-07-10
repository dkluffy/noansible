package core

import (
	"errors"
	//"noansible/mod"
	"noansible/mod"
	"noansible/target"
)

// windows 不支持Plugin
// func (taskm *TaskModule) LoadPlugin() error {
// 	var cmd_plus string

//
// 	if taskm.Shell == "" {
// 		taskm.Shell = cmd_plus
// 	} else {
// 		taskm.Shell += " " + cmd_plus
// 	}

// 	return nil
// }

type TaskModule struct {
	Name    string            `yaml:"name"`
	Shell   string            `yaml:"shell"`
	Include string            `yaml:"include"`
	Async   bool              `yaml:"async"`
	Plugin  map[string]string `yaml:"plugin"`
}

func (tsk *TaskModule) Shoot(t target.Target, tasklogs TaskLogs) error {
	var err error
	var result target.TargetStd
	if modName, ok := tsk.Plugin["mod"]; ok {

		modfunc, _ := mod.ModList[modName].(mod.ModCaller)

		args, ok := tsk.Plugin["arg"]
		if ok {
			args = tsk.Plugin["arg"]
		}
		result, err = modfunc.Run(t, args)
	} else {
		result, err = t.Execute(tsk.Shell)
	}

	tasklogs.Logger(tsk, result, err)
	if result.StdErr != "" {
		err = errors.New(result.StdErr)
	}
	return err
}

//for log
type TaskLog struct {
	IsFailed     bool
	TaskName     string
	ReturnValues map[string]interface{}
	ErrorInfo    string
}

type TaskLogs []TaskLog

func (tsklogs *TaskLogs) Logger(tsk *TaskModule, result target.TargetStd, err error) {
	var tlog TaskLog
	tlog.TaskName = tsk.Name

	if err != nil {
		tlog.IsFailed = true
		tlog.ErrorInfo = err.Error()
	} else if result.StdErr != "" {
		tlog.IsFailed = true
		tlog.ErrorInfo = string(result.StdErr)
	} else {
		tlog.IsFailed = false
		tlog.ReturnValues["StdOut"] = result.StdOut
	}
	*tsklogs = append(*tsklogs, tlog)
}
