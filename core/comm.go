package core

import (
	"fmt"
	"log"
	"noansible/target"
	"sync"
)

type PlayBookCommon struct {
	//在使用TAG 去yaml.Unmarshal的时候，struct的字段名首字母必须是大写的
	//否则，不会自动导入
	Hosts    string            `yaml:"hosts" json:"hosts"`
	Vars     map[string]string `yaml:"vars"`
	Username string            `yaml:"username"`
	tasklist []TaskModule
	hosts    []target.Hostinfo
}

var PlaybookVars map[string]string

type Playbook interface {
	Loader(filedir string, hostfile string)
	GetHead() PlayBookCommon
	//Player(hostlogs HostLogs)
}

var Gwaitgroup sync.WaitGroup

func Run(pb Playbook, hostlogs HostLogs, filedir string, hostfile string) {
	pb.Loader(filedir, hostfile)
	var pbc = pb.GetHead()
	for _, v := range pbc.hosts {
		Gwaitgroup.Add(1)
		var tasklogs TaskLogs
		go func(h target.Hostinfo) {
			h.Username = pbc.Username
			t := h.GetTarget()
			log.Println("**Connecting to: ", h.IPADDR)
			err := t.Connect()
			if err != nil {
				log.Println("  -- Can't connect to", h.IPADDR)
				msg := fmt.Sprintf("HostFailed@%v", h.IPADDR)
				tasklogs.SimpleLogger(msg, err)
			} else {
				log.Println("**OK, Connected to : ", h.IPADDR)
				for _, tk := range pbc.tasklist {
					err := tk.Shoot(t, &tasklogs)
					if err != nil {
						log.Println("  -- Task Failed: ", tk.Name, "@", h.IPADDR)
						if !tk.Ignore {
							break
						}
					}
				}
			}
			Gwaitgroup.Done()
			hostlogs[h.IPADDR] = tasklogs
		}(v)
	}
	Gwaitgroup.Wait()

}
