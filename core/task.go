package core

import (
	"errors"
	"log"

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

func (tsk *TaskModule) RunTask(t target.Target, tasklogs *TaskLogs) error {
	var err error
	var result target.TargetStd
	if modName, ok := tsk.Plugin["mod"]; ok {

		modinterface, ok := mod.ModList[modName]
		if !ok {
			err = errors.New("Mod not found, mod name: " + modName)
			tasklogs.Logger(tsk.Name, result, err)
			return err
		}
		modfunc, _ := modinterface.(mod.ModCaller)
		args := tsk.Plugin["args"]
		result, err = modfunc.Run(t, args)

	} else {
		result, err = t.Execute(tsk.Shell)
	}

	tasklogs.Logger(tsk.Name, result, err)
	if result.StdErr != "" {
		err = errors.New(result.StdErr)
	}
	return err
}

func (tsk *TaskModule) Shoot(t target.Target, tasklogs *TaskLogs) error {
	var err error
	log.Println("**Shooting Task: ", tsk.Name)
	if tsk.Async {
		go func() {
			err = tsk.RunTask(t, tasklogs)
			err = nil
		}()

	} else {
		return tsk.RunTask(t, tasklogs)
	}
	return err
}

//for log
type TaskLog struct {
	IsFailed     bool
	TaskName     string
	ReturnValues map[string]string
	ErrorInfo    string
}

type TaskLogs []TaskLog

func (tsklogs *TaskLogs) Logger(tskName string, result target.TargetStd, err error) {
	var tlog TaskLog
	tlog.ReturnValues = make(map[string]string)

	tlog.TaskName = tskName

	if err != nil {
		tlog.IsFailed = true
		tlog.ErrorInfo = err.Error()
	} else if result.StdErr != "" {
		tlog.IsFailed = true
		tlog.ErrorInfo = string(result.StdErr)
	} else {
		tlog.IsFailed = false
	}
	tlog.ReturnValues["StdOut"] = result.StdOut
	*tsklogs = append(*tsklogs, tlog)
}

func (tsklogs *TaskLogs) SimpleLogger(msg string, err error) {
	var tlog TaskLog
	tlog.TaskName = msg
	tlog.IsFailed = true
	tlog.ErrorInfo = err.Error()
	*tsklogs = append(*tsklogs, tlog)
}
