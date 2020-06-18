package core

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

//---yaml 实现
//TODO: func formatter
//这个第三方库，很不方便，不能兼容ansible的YML， 有空写个
type playbookYML struct {
	Playbook `yaml:",inline"`
	TasksYML []yaml.Node `yaml:"tasks"`
}

func (pbyml *playbookYML) Loader(filedir string) error {
	data, err := ioutil.ReadFile(filedir)
	if err != nil {
		log.Fatalf("unable to read playbook: %v", err)
		return err
	}
	err = yaml.Unmarshal(data, &pbyml)
	if err != nil {
		log.Fatalf("unable to read playbook: %v", err)
		return err
	}

	for _, v := range pbyml.TasksYML {
		var tmptask TaskModule //传址-必须放循环里，不然会搞笑
		terr := v.Decode(&tmptask)
		if terr != nil {
			log.Fatalf("Copy error task:%v ", terr)
			return terr
		} else {
			//TODO(3):在这里加载plugin -- tmptask.LoadPlugin()
			pbyml.tasks = append(pbyml.tasks, tmptask)
		}

	}

	return err
}

//TODO:json实现

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
