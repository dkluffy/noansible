package core

//---yaml 实现
import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sync"

	"noansible/target"

	"gopkg.in/yaml.v3"
)

//TODO: func formatter
//这个第三方库，很不方便，不能兼容ansible的YML， 有空写个

type PlaybookYML struct {
	PlaybookHead `yaml:",inline"`
	TasksYML     []yaml.Node `yaml:"tasks"`
	tasklist     []TaskModule
	hosts        []target.Hostinfo
}

var rootPath string

func (pbyml *PlaybookYML) Player(hostlogs HostLogs) {
	var wg sync.WaitGroup
	wg.Add(len(pbyml.hosts))
	for _, v := range pbyml.hosts {
		var tasklogs TaskLogs
		go func(h target.Hostinfo) {
			h.Username = pbyml.Username
			t := h.GetTarget()
			log.Println("**Connecting to: ", h.IPADDR)
			err := t.Connect()
			if err != nil {
				log.Println("  -- Can't connect to", h.IPADDR)
				msg := fmt.Sprintf("HostFailed@%v", h.IPADDR)
				tasklogs.SimpleLogger(msg, err)
			} else {
				log.Println("**OK, Connected to : ", h.IPADDR)
				for _, tk := range pbyml.tasklist {
					err := tk.Shoot(t, &tasklogs)
					if err != nil {
						log.Println("  -- Task Failed: ", tk.Name, "@", h.IPADDR)
						break
					}
				}
			}
			wg.Done()
			hostlogs[h.IPADDR] = tasklogs
		}(v)
	}
	wg.Wait()

}

func (pbyml *PlaybookYML) Loader(filedir string, hostfile string) {
	var err error
	rootPath = filepath.Dir(filedir)

	*pbyml, err = parseBook(filedir)
	if err != nil {
		log.Fatalf("Fail to Parse playbook\n Error:%v", err)
	}

	pbyml.tasklist, err = decodeTasks(pbyml.TasksYML)
	if err != nil {
		log.Fatalf("Fail to Decode Tasks\n Error:%v", err)
	}
	pbyml.hosts, err = ReadInventory(pbyml.Hosts, hostfile)
	if err != nil {
		log.Fatalf("Fail to Load inventory %v\n Error:%v", hostfile, err)
	}

	PlaybookVars = pbyml.Vars

	//render Vars
	for k, v := range PlaybookVars {
		nv, err := Render(v)
		if err != nil {
			log.Fatalf("Fail to Parse playbook\n Error:%v", err)
		} else {
			PlaybookVars[k] = nv
		}
	}
}

func parseBook(filedir string) (PlaybookYML, error) {
	var pbyml PlaybookYML
	data, err := ioutil.ReadFile(filedir)
	if err != nil {
		return pbyml, err
	}
	err = yaml.Unmarshal(data, &pbyml)
	if err != nil {
		return pbyml, err
	}

	return pbyml, err
}

func decodeTasks(taskNodes []yaml.Node) ([]TaskModule, error) {
	var tasks []TaskModule
	var terr error
	for _, v := range taskNodes {
		var tmptask TaskModule //传址-必须放循环里，不然会搞笑
		err := v.Decode(&tmptask)
		if err != nil {
			terr = err
			break
		}
		if tmptask.Include != "" { //被include的playbook只读取tasks，忽略其他
			pbyml, err := parseBook(filepath.Join(rootPath, tmptask.Include))
			if err != nil {
				terr = err
				break
			}
			tmptasks, err := decodeTasks(pbyml.TasksYML)
			if err != nil {
				terr = err
				break
			}
			tasks = append(tasks, tmptasks...)
			continue
		}
		tasks = append(tasks, tmptask)
	}

	return tasks, terr
}

//-------------------

func loadrawbook(filedir string) (map[interface{}]interface{}, error) {
	data, err := ioutil.ReadFile(filedir)
	book := make(map[interface{}]interface{})
	if err != nil {
		return book, err
	}

	err = yaml.Unmarshal(data, &book)
	//log.Fatalf("unable to read playbook: %v ---", err)

	return book, err

}
