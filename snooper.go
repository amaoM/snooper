package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	fp := `/proc/stat`
	flst := GetData(fp)
	time.Sleep(1 * time.Second)
	slst := GetData(fp)
	Calculate(flst, slst, os.Stdout)
}

func GetData(fp string) []int {
	f, err := OpenFile(fp)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	t, err := GetText(f)
	if err != nil {
		log.Fatal(err)
	}
	slst, err := SplitText(t)
	if err != nil {
		log.Fatal(err)
	}
	ilst, err := ParseStat(slst)
	if err != nil {
		log.Fatal(err)
	}
	return ilst
}

func OpenFile(fp string) (*os.File, error) {
	f, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func GetText(f *os.File) (string, error) {
	s := bufio.NewScanner(f)

	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}

	t := s.Text()
	if t == "" || !regexp.MustCompile(`^cpu  `).Match([]byte(t)) {
		return "", errors.New("Failed to get a text")
	}

	return t, nil
}

func SplitText(t string) ([]string, error) {
	lst := strings.Split(t, " ")
	if 12 > len(lst) {
		return nil, errors.New("Parse data is invalid data")
	}
	return lst, nil
}

func ParseStat(lst []string) ([]int, error) {
	var tlst []int

	for i := 0; i < len(lst); i++ {
		if regexp.MustCompile(`[0-9]`).Match([]byte(lst[i])) {
			t, err := strconv.Atoi(lst[i])
			if err != nil {
				return nil, err
			}
			tlst = append(tlst, t)
		}
	}

	return tlst, nil
}

func Calculate(flst []int, slst []int, w io.Writer) {
	tlst := make([]int, len(flst))
	sum := 0

	for i := 0; i < len(flst); i++ {
		tlst[i] = slst[i] - flst[i]
		sum += tlst[i]
	}

	items := []string{"user", "nice", "system", "idle", "iowait", "irq", "softirq", "steal", " ?"}
	for ii := 0; ii < len(items); ii++ {
		fmt.Fprintf(w, "%-8s", items[ii])
		fmt.Fprintf(w, "%3s", strconv.Itoa(tlst[ii]*100/sum))
		fmt.Fprintln(w, " %")
	}
}
