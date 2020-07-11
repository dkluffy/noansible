package mod

import (
	"errors"
	"fmt"
	"noansible/target"
	"strings"
)

type FileMod struct {
	src string
	dst string
}

func (f *FileMod) Run(t target.Target, args string) (target.TargetStd, error) {

	args = strings.ReplaceAll(args, " ", "")
	arg := strings.Split(args, ",")

	if len(arg) < 2 {
		var t target.TargetStd
		err := errors.New("fileModError:too few args,or format error!")
		return t, err
	}
	f.src = arg[0]
	f.dst = arg[1]

	if f.src[:1] == "@" {
		//copy on remote only
		cmd := fmt.Sprintf("cp -rf %s %s", f.src[1:], f.dst)
		return t.Execute(cmd)
	} else {
		return t.Copy(f.src, f.dst, target.BUFFERSIZE)
	}
}

func (f *FileMod) NewMod(args map[string]string) {
	f.src = args["src"]
	f.dst = args["dst"]
}
