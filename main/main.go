package main

import (
	"flag"
	"fmt"

	//"fmt"
	"noansible/core"
)

// 主调用
func play(pb core.Playbook, pbfile string, hostsfile string, tklogs *core.TaskLogs) {
	pb.Loader(pbfile, hostsfile)
	pb.Player(tklogs)
}

func main() {
	fmt.Println("Noansible @version=", "0.1")

	//command args
	hostfile := flag.String("i", "inventory.yml", "Inventory file dir")
	playbookfile := flag.String("p", "main.yml", "Inventory file dir")

	flag.Parse()

	//load yml playbook
	var pbyml core.PlaybookYML
	var tklogs core.TaskLogs

	play(&pbyml, *playbookfile, *hostfile, &tklogs)
	fmt.Println("----tklogs:", tklogs)
}
