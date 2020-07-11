package mod

import (
	"noansible/target"
	"time"
)

type SyncTimeMod struct {
}

func (dt *SyncTimeMod) Run(t target.Target, args string) (target.TargetStd, error) {
	localtime := time.Now().Format(time.UnixDate)
	return t.Execute("date -s \"" + localtime + "\"")
}
