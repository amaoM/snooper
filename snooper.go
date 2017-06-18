package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	fssh := flag.String("ssh", "", "ssh connection")
	fftp := flag.String("sftp", "", "sftp connection")

	flag.Parse()

	if *fssh != "" && *fftp != "" {
		log.Fatal("You cannot specify ssh and sftp together")
	}

	if *fssh != "" {
		user, host, err := SplitUserHost(*fssh)

		if err != nil {
			log.Fatal(err)
		}

		err = Ssh(host, user)

		if err != nil {
			log.Fatal(err)
		}

	} else if *fftp != "" {
		user, host, err := SplitUserHost(*fftp)

		if err != nil {
			log.Fatal(err)
		}

		fp := "/proc/stat"
		ofp := os.Getenv("HOME") + "/stac"
		err = Sftp(host, user, fp, ofp)
		if err != nil {
			log.Fatal(err)
		}

		flst := GetData(ofp)
		time.Sleep(1 * time.Second)

		err = Sftp(host, user, fp, ofp)
		if err != nil {
			log.Fatal(err)
		}

		slst := GetData(ofp)
		Calculate(flst, slst, os.Stdout)

	} else {
		fp := `/proc/stat`
		Local(fp)
	}
}

func SplitUserHost(arg string) (string, string, error) {
	lst := strings.Split(arg, "@")
	if len(lst) != 2 {
		return "", "", errors.New("Invalid arguments")
	}
	return lst[0], lst[1], nil
}
