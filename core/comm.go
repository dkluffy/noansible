package core

import "sync"

type PlaybookHead struct {
	//在使用TAG 去yaml.Unmarshal的时候，struct的字段名首字母必须是大写的
	//否则，不会自动导入
	Hosts    string            `yaml:"hosts" json:"hosts"`
	Vars     map[string]string `yaml:"vars"`
	Username string            `yaml:"username"`
}

var PlaybookVars map[string]string

type Playbook interface {
	Loader(filedir string, hostfile string)
	Player(hostlogs HostLogs)
}

var Gwaitgroup sync.WaitGroup
