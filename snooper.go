package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	log.Println("started")
	var wg sync.WaitGroup
	for _, ip := range []string{"192.168.20.20", "192.168.70.71"} {
		wg.Add(1)
		go getCpuUsage(ip, &wg)
	}
	wg.Wait()
	log.Println("finished")
}

func getCpuUsage(ip string, wg *sync.WaitGroup) {
	h := new(Host)
	h.ip = ip
	h.port = "22"
	h.user = "vagrant"

	fs := new(Stat)
	fs.path = "/proc/stat"
	ss := new(Stat)
	ss.path = "/proc/stat"

	for _, s := range []*Stat{fs, ss} {
		err := h.Connect(s)
		if err != nil {
			log.Fatal(err)
		}
		defer h.client.Close()
		defer h.session.Close()

		// To reduce the influence of a ssh connection processing
		time.Sleep(1 * time.Second)

		err = h.execCmd(s)
		if err != nil {
			log.Fatal(err)
		}

		bstr, err := s.changeBufferToString()
		if err != nil {
			log.Fatal(err)
		}

		err = s.splitStatString(bstr)
		if err != nil {
			log.Fatal(err)
		}

		slst, err := s.splitCpuUsageTimeString()
		if err != nil {
			log.Fatal(err)
		}

		err = s.getCpuUsageTime(slst)
		if err != nil {
			log.Fatal(err)
		}
	}

	calculate(fs.cpuUsageTimeList, ss.cpuUsageTimeList, os.Stdout, h.ip)

	wg.Done()
}

func calculate(flst []int, slst []int, w io.Writer, ip string) {
	tlst := make([]int, len(flst))
	sum := 0

	for i := 0; i < len(flst); i++ {
		tlst[i] = slst[i] - flst[i]
		sum += tlst[i]
	}

	items := []string{"user", "nice", "system", "idle", "iowait", "irq", "softirq", "steal", "guest"}

	fmt.Fprintln(w, "+++++ "+ip+" +++++")
	for ii := 0; ii < len(items); ii++ {
		fmt.Fprintf(w, "%-8s", items[ii])
		fmt.Fprintf(w, "%3s", strconv.Itoa(tlst[ii]*100/sum))
		fmt.Fprintln(w, " %")
	}
	fmt.Fprintln(w, "+++++++++++++++++++++++++")
}
