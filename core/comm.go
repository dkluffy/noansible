package core

import (
	"bytes"
)

type hostinfo struct {
	ipaddr   string
	port     int
	username string
	password string
}

type ExecResult struct {
	stdOut, stdErr bytes.Buffer
}

type TaskModule struct {
	Name    string `yaml:"name"`
	Shell   string `yaml:"shell"`
	Include string `yaml:"include"`
	//Plugin  interface{} `yaml:"plugin"`
	Plugin map[interface{}]interface{} `yaml:"plugin"`
	//returnValue interface{} //TODO
}

type Playbook struct {
	//在使用TAG 去yaml.Unmarshal的时候，struct的字段名首字母必须是大写的
	//否则，不会自动导入
	Hosts    string            `yaml:"hosts" json:"hosts"`
	Vars     map[string]string `yaml:"vars"`
	Username string            `yaml:"username"`
	tasks    []TaskModule      //- 这个是为了json等通用支持
}
