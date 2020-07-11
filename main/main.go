package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"noansible/core"
	"noansible/target"
	"os"
)

// 主调用
func play(pb core.Playbook, pbfile string, hostsfile string, hostlogs core.HostLogs) {
	pb.Loader(pbfile, hostsfile)
	pb.Player(hostlogs)
}

func main() {
	fmt.Println("Noansible @version=", "1.0")

	//command args
	hostfile := flag.String("i", "inventory.yml", "Inventory file dir")
	playbookfile := flag.String("p", "main.yml", "Inventory file dir")
	logfiledir := flag.String("log", "output.log", "Log file dir")
	buffersize := flag.Int("bs", 1024, "SCP buffer size")

	target.BUFFERSIZE = *buffersize

	flag.Parse()

	f, err := os.OpenFile(*logfiledir, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	// 定义多个写入器
	writers := []io.Writer{
		f,
		os.Stdout}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	log.SetOutput(fileAndStdoutWriter)

	//load yml playbook
	var pbyml core.PlaybookYML
	hostlogs := make(core.HostLogs)
	play(&pbyml, *playbookfile, *hostfile, hostlogs)
	hostlogs.Printer()
}
