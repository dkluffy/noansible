package core

import (
	"fmt"
	//"reflect"
	"testing"
)

func TestSSHTarget_Execute(t *testing.T) {
	// type args struct {
	// 	cmd string
	// }
	host1 := &SSHTarget{
		ipaddr:   "[fd81::169]",
		port:     22,
		username: "root",
		password: "1",
	}

	if err1 := host1.Connect(); err1 != nil {
		t.Errorf("host1-error = %v", err1)
		return
	}

	t.Run("host1", func(t *testing.T) {
		ert, err := host1.Execute("echo host1 && date >>/tmp/h.txt")
		ert, err = host1.Execute("echo host1 && echo \"--$(date)\" >>/tmp/h.txt")
		fmt.Println(ert)
		if err != nil {
			t.Errorf("result= %v ,host1-error = %v", ert, err)
			return
		}

	})

	// tests := []struct {
	// 	name    string
	// 	ss      *SSHTarget
	// 	args    args
	// 	want    ExecResult
	// 	wantErr bool
	// }{
	// 	{
	// 		name: "host1",
	// 		ss: host1,
	// 		args: { cmd: "echo host1 && echo host1 >/tmp/h.txt"},
	// 	},
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		got, err := tt.ss.Execute(tt.args.cmd)
	// 		if (err != nil) != tt.wantErr {
	// 			t.Errorf("SSHTarget.Execute() error = %v, wantErr %v", err, tt.wantErr)
	// 			return
	// 		}
	// 		if !reflect.DeepEqual(got, tt.want) {
	// 			t.Errorf("SSHTarget.Execute() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}
