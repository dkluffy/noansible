package core

type PlaybookHead struct {
	//在使用TAG 去yaml.Unmarshal的时候，struct的字段名首字母必须是大写的
	//否则，不会自动导入
	Hosts    string            `yaml:"hosts" json:"hosts"`
	Vars     map[string]string `yaml:"vars"`
	Username string            `yaml:"username"`
}

// type Task interface {
// 	Runner(t Target, tasklog *[]string) error
// 	//TODO:AsyncRunner(t Target, tasklog *[]string)
// 	Logger() (taskLogs *[]TaskLog, result TargetStd, err error)
// }

type Playebook interface {
	Loader(filedir string) error
	//Runner()
	//Reporter()
}
