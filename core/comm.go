package core

type Target interface {
	Connect() error
	Close()
	Execute(cmd string) (interface{}, error)
}
type hostinfo struct {
	ipaddr   string
	port     int
	username string
	password string
}

type TaskCooker interface {
}

type TaskModule struct {
	Name    string `yaml:"name"`
	Shell   string `yaml:"shell"`
	Include string `yaml:"include"`
	//Plugin  interface{} `yaml:"plugin"`
	Plugin map[interface{}]interface{} `yaml:"plugin"`
	//returnValue interface{} //TODO:返回值
}

type player interface {
	Loader(filedir string) error
	//Runner()
	//Reporter()
}
type Playbook struct {
	//在使用TAG 去yaml.Unmarshal的时候，struct的字段名首字母必须是大写的
	//否则，不会自动导入
	Hosts    string            `yaml:"hosts" json:"hosts"`
	Vars     map[string]string `yaml:"vars"`
	Username string            `yaml:"username"`
	tasks    []TaskModule      //- 这个是为了json等通用支持
}
