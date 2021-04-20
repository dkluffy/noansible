package target

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"

	"net"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type Hostinfo struct {
	IPADDR     string
	Port       int
	Username   string
	password   string
	privatekey string //key的路径
	//connectMethod interface{}
}

func (hinfo *Hostinfo) NewHost(item string) {
	hinfo.Port = 22
	item = strings.ReplaceAll(item, " ", "")
	vsarr := strings.Split(item, ",")
	if len(vsarr) >= 3 && vsarr[2] != "" {
		hinfo.Port, _ = strconv.Atoi(vsarr[2])
	}
	if strings.Contains(vsarr[0], ":") {
		hinfo.IPADDR = "[" + vsarr[0] + "]"
	} else {
		hinfo.IPADDR = vsarr[0]
	}

	hinfo.password = vsarr[1]
}

func (hinfo *Hostinfo) GetTarget() Target {
	var t Target
	if hinfo.Port == 22 {
		ht := SSHTarget{
			Hostinfo: *hinfo,
		}
		t = &ht
	}
	return t
}

type TargetStd struct {
	StdOut, StdErr string
}

type Target interface {
	Connect() error
	//Close()
	Execute(cmd string) (TargetStd, error)
	Copy(src string, dst string, buffersize int) (TargetStd, error)
}

//---ssh实现

type SSHTarget struct {
	Hostinfo
	client *ssh.Client
}

func (ss *SSHTarget) Connect() error {

	//var hostKey ssh.PublicKey
	hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}
	var err error

	addr := fmt.Sprintf("%s:%d", ss.IPADDR, ss.Port)

	config := &ssh.ClientConfig{
		User: ss.Username,
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
			//log.Fatalf("unable to read private key: %v", err)
			return err
		}
		// Create the Signer for this private key.
		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			//log.Fatalf("unable to parse private key: %v", err)
			return err
		}
		config.Auth = append(config.Auth, ssh.PublicKeys(signer))
	}

	ss.client, err = ssh.Dial("tcp", addr, config)
	if err != nil {
		//log.Fatal("Failed to dial: ", err)

		return err
	}

	return nil
}

func (ss *SSHTarget) Execute(cmd string) (TargetStd, error) {
	var result TargetStd

	var stdOut, stdErr bytes.Buffer

	session, err := ss.client.NewSession()
	if err != nil {
		return result, err
	}
	defer session.Close()

	session.Stderr = &stdErr
	session.Stdout = &stdOut

	if err = session.Run(cmd); err != nil {
		//log.Fatal("Failed to run: " + err.Error())
		return result, err
	}
	result.StdErr = stdErr.String()
	result.StdOut = stdOut.String()
	return result, err
}

//TODO: 更好的RESULT
func (ss *SSHTarget) Copy(src string, dst string, buffersize int) (TargetStd, error) {
	var result TargetStd
	//var stdOut, stdErr bytes.Buffer

	//create sftp client session from ssh client
	sftpc, err := sftp.NewClient(ss.client)
	if err != nil {
		return result, err
	}
	defer sftpc.Close()

	switch src[len(src)-2:] {
	case "/*":
		err = scpDir(sftpc, src, dst, buffersize)

	default:
		err = scpFile(sftpc, src, dst, buffersize)

	}

	return result, err

}

//TODO-1:fectch远程文件

func scpDir(sftpc *sftp.Client, src string, dst string, buffersize int) error {
	//src: ../abc/*  一定要这个格式以明确，以/*号结尾
	//dst: /root/abcd  必须指定目录
	if src[len(src)-1:] != "*" {
		return errors.New("scpDir Failed: src is not end with * , such as ./ab/cd/* ")
	}

	var srcFileList []string
	src = src[:len(src)-1]
	err := FindAllFiles(src, &srcFileList)
	if err != nil {
		return err
	}

	dstFileList := RemovePrefix(srcFileList, src)

	err = sftpc.MkdirAll(dst)
	if err != nil {
		return err
	}

	//建立目录树
	re := regexp.MustCompile(`\\`)
	newps := SSHPS
	for i := 0; i < len(dstFileList); i++ {
		dstFileList[i] = re.ReplaceAllString(dstFileList[i], newps)
		if dstFileList[i][len(dstFileList[i])-1:] == newps {
			err = sftpc.MkdirAll(dst + newps + dstFileList[i])
			if err != nil {
				return err
			}
		}
	}
	//复制文件
	for i := 0; i < len(dstFileList); i++ {

		if dstFileList[i][len(dstFileList[i])-1:] == newps {
			continue
		}
		//fmt.Println("****Debug: Copy file to remote ", srcFileList[i], "dst:", dst+newps+dstFileList[i])
		err = scpFile(sftpc, srcFileList[i], dst+newps+dstFileList[i], buffersize)
		if err != nil {
			return err
		}
	}
	return err

}

func scpFile(sftpc *sftp.Client, src string, dst string, buffersize int) error {
	//create sftp client session from ssh client
	// sftpc, err := sftp.NewClient(client)
	// if err != nil {
	// 	return err
	// }
	// defer sftpc.Close()

	srcFile, err := os.Open(src)

	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFileName := dst
	if dst[len(dst)-1:] == "." {
		dstFileName = filepath.Join(dst, filepath.Base(src))
	}

	dstFile, err := sftpc.Create(dstFileName)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	buf := make([]byte, buffersize)
	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		_, err := dstFile.Write(buf[0:n])
		if err != nil {
			return err
		}
	}
	return nil

}
