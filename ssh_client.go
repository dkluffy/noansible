package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os/exec"
	"strconv"
	str "strings"

	"golang.org/x/crypto/ssh"
)

func runCmd() {

	var stdOut, stdErr bytes.Buffer

	cmd := exec.Command("ssh", "root@fd81::169", "if [ -d liujx/project ];then echo 0;else echo \"123aaa\";fi")
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	if err := cmd.Run(); err != nil {
		fmt.Printf("cmd exec failed: %s : %s", fmt.Sprint(err), stdErr.String())
	}

	fmt.Print(stdOut.String())
	ret, err := strconv.Atoi(str.Replace(stdOut.String(), "\n", "", -1))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d, %s\n", ret, stdErr.String())
} //key登录

//
func SSHConnect(user, password, host string, port int) (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}

	clientConfig = &ssh.ClientConfig{
		User: user,
		Auth: auth,
		// Timeout:             30 * time.Second,
		HostKeyCallback: hostKeyCallbk,
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil
}

func runSsh() {

	var stdOut, stdErr bytes.Buffer

	session, err := SSHConnect("root", "1", "192.168.0.116", 22)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	session.Stdout = &stdOut
	session.Stderr = &stdErr

	session.Run("echo '123a'")
	fmt.Println(stdOut.String())
	// ret, err := strconv.Atoi(str.Replace(stdOut.String(), "\n", "", -1))
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("%d, %s\n", ret, stdErr.String())

}

func main() {
	runSsh()
}
