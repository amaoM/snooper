package main

import (
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ssh"
)

func Connect(user string, host string) (*ssh.Client, error) {
	port := "22"

	buf, err := ioutil.ReadFile(os.Getenv("HOME") + "/.ssh/id_rsa")
	if err != nil {
		return nil, err
	}

	key, err := ssh.ParsePrivateKey(buf)
	if err != nil {
		return nil, err
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", host+":"+port, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
