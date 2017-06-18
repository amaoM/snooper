package main

import (
	"bytes"
	"fmt"
	"log"
)

func Ssh(host string, user string) error {
	conn, err := Connect(user, host)
	if err != nil {
		return err
	}

	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b

	if err := session.Run("cat /proc/stat"); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}

	fmt.Println(b.String())

	return nil
}
