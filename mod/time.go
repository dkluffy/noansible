package mod

import (
	"noansible/target"
)

type SyncTimeMod struct {
}

func (dt *SyncTimeMod) Run(t target.Target,args string) (target.TargetStd, error) {
	return t.Execute("date >/tmp/synctime.txt")
}
