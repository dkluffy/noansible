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

var Versiondate = "2020/07/24"
var Version = "2.0a"
var versiontstr = fmt.Sprintf("\nNoansible@%v version: %v (By dkluffy@gmail.com)", Versiondate, Version)

// 主调用
func play(pb core.Playbook, pbfile string, hostsfile string, hostlogs core.HostLogs) {
	pb.Loader(pbfile, hostsfile)
	pb.Player(hostlogs)
}

func usage() {

	fmt.Fprint(os.Stderr, versiontstr)
	fmt.Fprint(os.Stderr, "https://github.com/dkluffy/noansible")
	fmt.Fprint(os.Stderr, "\n------\n")
	fmt.Fprintf(os.Stderr, `
Usage: noansible [-h] [-i inventoryfile] [-p playbookfile] [-bs buffersize] [-log logfile]

Options:
`)
	flag.PrintDefaults()
}

func main() {

	//command args
	help := flag.Bool("h", false, "print this help")
	hostfile := flag.String("i", "inventory.yml", "Inventory file dir")
	playbookfile := flag.String("p", "main.yml", "Inventory file dir")
	logfiledir := flag.String("log", "output.log", "Log file dir")
	buffersize := flag.Int("bs", 1024, "SCP buffer size")

	flag.Usage = usage

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Fprint(os.Stderr, versiontstr)

	target.BUFFERSIZE = *buffersize
	f, err := os.OpenFile(*logfiledir, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
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
