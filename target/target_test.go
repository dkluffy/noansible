package target

import (
	"fmt"
	//"reflect"
	"testing"

	"github.com/pkg/sftp"
)

func TestCopy(t *testing.T) {
	host1 := &SSHTarget{
		Hostinfo: Hostinfo{
			IPADDR:   "[fd81::169]",
			Port:     22,
			Username: "root",
			password: "1",
		},
	}

	if err1 := host1.Connect(); err1 != nil {
		t.Errorf("host1-error = %v", err1)
		return
	}

	t.Run("host1", func(t *testing.T) {

		//err := scpFile(sftpc, "../files", "/tmp/02", 1024)
		//err = sftpc.MkdirAll("/tmp/01/02")
		_, err := host1.Copy("../files/*", "/tmp/tttt", 1024)
		t.Errorf("result= %v ", err)

		if err != nil {
			t.Errorf("result= %v ", err)
			return
		}

	})

}

func TestScpFile(t *testing.T) {
	host1 := &SSHTarget{
		Hostinfo: Hostinfo{
			IPADDR:   "[fd81::169]",
			Port:     22,
			Username: "root",
			password: "1",
		},
	}

	if err1 := host1.Connect(); err1 != nil {
		t.Errorf("host1-error = %v", err1)
		return
	}
	//create sftp client session from ssh client
	sftpc, err := sftp.NewClient(host1.client)
	if err != nil {
		t.Errorf("host1-error = %v", err)
		return
	}
	defer sftpc.Close()

	t.Run("host1", func(t *testing.T) {

		//err := scpFile(sftpc, "../files", "/tmp/02", 1024)
		//err = sftpc.MkdirAll("/tmp/01/02")
		err = scpDir(sftpc, "../files/*", "/tmp", 1024)
		t.Errorf("result= %v ", err)

		if err != nil {
			t.Errorf("result= %v ", err)
			return
		}

	})

}

func TestSCP(t *testing.T) {
	// type args struct {
	// 	cmd string
	// }
	host1 := &SSHTarget{
		Hostinfo: Hostinfo{
			IPADDR:   "[fd81::169]",
			Port:     22,
			Username: "root",
			password: "1",
		},
	}

	// var host1 SSHTarget
	// host1.ipaddr = "[fd81::169]"

	if err1 := host1.Connect(); err1 != nil {
		t.Errorf("host1-error = %v", err1)
		return
	}

	t.Run("host1", func(t *testing.T) {
		session, err := host1.client.NewSession()
		if err != nil {
			t.Errorf("Errr: %v", err)
		}
		defer session.Close()
		go func() {
			w, _ := session.StdinPipe()
			defer w.Close()
			content := "123456789\n"
			fmt.Fprintln(w, "D0755", 0, "testdir") // mkdir
			fmt.Fprintln(w, "C0644", len(content), "testfile1")
			fmt.Fprint(w, content)
			fmt.Fprint(w, "\x00") // transfer end with \x00
			fmt.Fprintln(w, "C0644", len(content), "testfile2")
			fmt.Fprint(w, content)
			fmt.Fprint(w, "\x00")
		}()
		if err := session.Run("/usr/bin/scp -tr /tmp/."); err != nil {
			t.Errorf("Errr: %v", err)
		}

	})
}

func TestSSHTarget_Execute(t *testing.T) {
	// type args struct {
	// 	cmd string
	// }
	host1 := &SSHTarget{
		Hostinfo: Hostinfo{
			IPADDR:   "[fd81::169]",
			Port:     22,
			Username: "root",
			password: "1",
		},
	}

	// var host1 SSHTarget
	// host1.ipaddr = "[fd81::169]"

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
