package mod

import (
	"noansible/target"
)

//这样在调用的时候 ModList["file"].Run(...)
type ModCaller interface {
	Run(t target.Target, args string) (target.TargetStd, error)
}

var (
	ModList = map[string]ModCaller{
		"file": &FileMod{},
		"time": &SyncTimeMod{},
	}
)

//---regexp

//var reMultispace = regexp.MustCompile(` {2,}`)
//args = reMultispace.ReplaceAllString(args, " ")
