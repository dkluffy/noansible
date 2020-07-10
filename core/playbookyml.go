package core

//---yaml 实现
import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

//TODO: func formatter
//这个第三方库，很不方便，不能兼容ansible的YML， 有空写个

type PlaybookYML struct {
	PlaybookHead `yaml:",inline"`
	TasksYML     []yaml.Node `yaml:"tasks"`
	tasklist     []TaskModule
}

func (pbyml *PlaybookYML) Loader(filedir string) error {
	var err error

	*pbyml, err = parseBook(filedir)
	if err != nil {
		return err
	}

	pbyml.tasklist, err = decodeTasks(pbyml.TasksYML)
	return err
}

func parseBook(filedir string) (PlaybookYML, error) {
	var pbyml PlaybookYML
	data, err := ioutil.ReadFile(filedir)
	if err != nil {
		log.Fatalf("unable to parse playbook: %v", err)
		return pbyml, err
	}
	err = yaml.Unmarshal(data, &pbyml)
	if err != nil {
		log.Fatalf("unable to parse playbook: %v", err)
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
			log.Fatalf("Can't decode task:%v ", err)
			terr = err
			break
		}
		if tmptask.Include != "" { //被include的playbook只读取tasks，忽略其他
			pbyml, err := parseBook(tmptask.Include)
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
	if err != nil {
		log.Fatalf("unable to read playbook: %v", err)
	}

	book := make(map[interface{}]interface{})
	err = yaml.Unmarshal(data, &book)
	//log.Fatalf("unable to read playbook: %v ---", err)

	return book, err

}

// func parsemap(bookmap *map[interface{}]interface{}) (playbook,error){

// 	for k,v := range *bookmap {
// 		switch k {
// 		case "vars":

// 		}
// 	}
// }
