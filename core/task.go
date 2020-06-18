package core

//TODO(1):写task接口，先实现命令行批量执行

func (taskm *TaskModule) LoadPlugin() error {
	var cmd_plus string

	//TODO(2)解析plugin，运行后得到 cmd_plus
	if taskm.Shell == "" {
		taskm.Shell = cmd_plus
	} else {
		taskm.Shell += " " + cmd_plus
	}

	return nil
}
