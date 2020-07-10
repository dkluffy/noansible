package main

import (
	"fmt"
	"plugin"
)

func RunPlugin(libdir string) error {
	//打开动态库
	pdll, err := plugin.Open(libdir)
	if err != nil {
		fmt.Println("Load error", err)
		return err
	}
	PrintNow, err := pdll.Lookup("PrintNow")
	if err != nil {
		return err
	}
	t := PrintNow.(func() string)()
	fmt.Println(t)

	return err

}

func main() {

	err := RunPlugin("time.so")
	if err != nil {
		fmt.Println("--main:", err)
	}

}
