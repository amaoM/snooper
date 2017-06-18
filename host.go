package main

import (
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ssh"
)

type Host struct {
	ip      string
	user    string
	port    string
	client  *ssh.Client
	session *ssh.Session
}

func (h *Host) Connect(stat *Stat) error {
	err := h.Ssh()
	if err != nil {
		return err
	}

	h.session, err = h.client.NewSession()
	if err != nil {
		return err
	}

	return nil
}

func (h *Host) Ssh() error {
	buf, err := ioutil.ReadFile(os.Getenv("HOME") + "/.ssh/id_rsa")
	if err != nil {
		return err
	}

	key, err := ssh.ParsePrivateKey(buf)
	if err != nil {
		return err
	}

	config := &ssh.ClientConfig{
		User: h.user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	h.client, err = ssh.Dial("tcp", h.ip+":"+h.port, config)
	if err != nil {
		return err
	}

	return nil
}

func (h *Host) execCmd(stat *Stat) error {
	h.session.Stdout = &stat.buffer
	if err := h.session.Run("cat " + stat.path); err != nil {
		return err
	}
	return nil
}
