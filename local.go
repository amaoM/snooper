package main

import (
	"os"
	"time"
)

func Local(fp string) {
	flst := GetData(fp)
	time.Sleep(1 * time.Second)
	slst := GetData(fp)
	Calculate(flst, slst, os.Stdout)
}
