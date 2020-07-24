package core

import (
	"errors"
	"fmt"
	"log"
	"strings"

	//"noansible/mod"
	"noansible/extra"
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
	If      string            `yaml:"if"`
	Return  string            `yaml:"return"`
}

//运行任务
func (tsk *TaskModule) runTask(t target.Target, tasklogs *TaskLogs) error {
	//--异步的任务，不能使用字段`Return`

	var err error
	var result target.TargetStd

	//前置条件检查,如果失败直接跳过这个TASK
	if len(tsk.If) > 0 {
		tsk.If, err = Render(tsk.If)
		if err != nil {
			tasklogs.Logger(tsk.Name, result, err)
			return err
		}
		if !extra.IfTester(tsk.If) {
			tasklogs.Logger(tsk.Name, result, err)
			return nil
		}

	}

	//run plugin
	//只执行PLUGIN，shell会被忽略
	//没有PLUGIN，才会执行SHELL
	if modName, ok := tsk.Plugin["mod"]; ok {

		modinterface, ok := mod.ModList[modName]
		if !ok {
			err = errors.New("Mod not found, mod name: " + modName)
			tasklogs.Logger(tsk.Name, result, err)
			return err
		}
		modfunc, _ := modinterface.(mod.ModCaller)
		args := tsk.Plugin["args"]

		//模板渲染
		args, err = Render(args)
		if err != nil {
			tasklogs.Logger(tsk.Name, result, err)
			return err
		}

		result, err = modfunc.Run(t, args)

	} else {
		//模板渲染
		tsk.Shell, err = Render(tsk.Shell)
		if err != nil {
			tasklogs.Logger(tsk.Name, result, err)
			return err
		}
		result, err = t.Execute(tsk.Shell)
	}

	tasklogs.Logger(tsk.Name, result, err)
	if result.StdErr != "" {
		err = errors.New(result.StdErr)
	}

	tsk.updateReturnVars(result)
	return err
}

//处理Return字段
func (tsk *TaskModule) updateReturnVars(result target.TargetStd) {
	if tsk.Return != "" {
		v := strings.TrimSuffix(result.StdOut, "\n")
		v = strings.TrimSuffix(v, "\r")
		PlaybookVars[tsk.Return] = v
	}
}

//调用 运行任务 runTask
func (tsk *TaskModule) Shoot(t target.Target, tasklogs *TaskLogs) error {
	var err error
	log.Println("**Shooting Task: ", tsk.Name)
	if tsk.Async {

		//必须要这样写，不然会 调用子协程的时候，tsk指针会移动到下一个
		var tsktmp TaskModule
		Gwaitgroup.Add(1)
		tsktmp = *tsk
		var result target.TargetStd
		tasklogs.Logger(tsktmp.Name, result, err)
		go func(tsk *TaskModule) {
			var tasklogstmp TaskLogs //必须，不然会污染全局
			defer Gwaitgroup.Done()
			err1 := tsk.runTask(t, &tasklogstmp)
			if err1 != nil {
				log.Println("  -- Async Task Error:", tsk, err)
			}

		}(&tsktmp)

	} else {
		return tsk.runTask(t, tasklogs)
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
type HostLogs map[string]TaskLogs

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

func (tsklogs *TaskLogs) Printer() int {
	totalfailed := 0

	pre_str := "OK"
	for _, v := range *tsklogs {
		if v.IsFailed {
			pre_str = "**Failed"
			totalfailed += 1
		}
		errstr := ""
		if v.ErrorInfo != "" {
			errstr = fmt.Sprintf("\nError:%s", v.ErrorInfo)
		}

		msg := fmt.Sprintf("%s [ %s ] %s %s", pre_str, v.TaskName, v.ReturnValues, errstr)
		log.Println(msg)
	}
	return totalfailed
}

func (hlogs *HostLogs) Printer() {
	totalfailed := 0
	for k, v := range *hlogs {
		log.Println("===================")
		log.Println(k)
		log.Println("===================")
		totalfailed += v.Printer()
	}
	log.Println("")
	log.Println("**************************************")
	log.Println(fmt.Sprintf("Total Failed: %v / %v", totalfailed, len(*hlogs)))
}
