package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

type Target interface {
	Connect() error
	Close()
	Execute(cmd string) (ExecResult, error)
}

//---

type SSHTarget struct {
	ipaddr     string
	port       int
	username   string
	password   string
	privatekey string //key的路径
	client     *ssh.Client
}

func (ss *SSHTarget) Connect() error {
	//var hostKey ssh.PublicKey
	hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}
	var err error

	addr := fmt.Sprintf("%s:%d", ss.ipaddr, ss.port)
	config := &ssh.ClientConfig{
		User: ss.username,
		//HostKeyCallback: ssh.FixedHostKey(hostKey),
		HostKeyCallback: hostKeyCallbk,
	}

	if ss.password != "" {
		config.Auth = append(config.Auth, ssh.Password(ss.password))

	} else if ss.privatekey != "" {
		//TODO: test key auth
		//read key
		key, err := ioutil.ReadFile(ss.privatekey)
		if err != nil {
			log.Fatalf("unable to read private key: %v", err)
			return err
		}
		// Create the Signer for this private key.
		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			log.Fatalf("unable to parse private key: %v", err)
			return err
		}
		config.Auth = append(config.Auth, ssh.PublicKeys(signer))
	}

	ss.client, err = ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
		return err
	}
	return nil
}

func (ss *SSHTarget) Execute(cmd string) (ExecResult, error) {
	var exec_result ExecResult

	session, err := ss.client.NewSession()
	if err != nil {
		return exec_result, err
	}
	defer session.Close()

	session.Stdout = &exec_result.stdOut
	session.Stderr = &exec_result.stdErr
	if err = session.Run(cmd); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
	return exec_result, err
}
