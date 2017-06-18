package main

import (
	"bytes"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Stat struct {
	path             string
	buffer           bytes.Buffer
	cpuUsageTimeText string
	cpuUsageTimeList []int
}

func (s *Stat) changeBufferToString() (string, error) {
	bstr := s.buffer.String()
	if bstr == "" {
		return "", errors.New("The Buffer is empty")
	}
	return bstr, nil
}

func (s *Stat) splitStatString(bstr string) error {
	tlst := strings.Split(bstr, "\n")
	if len(tlst) == 0 {
		return errors.New("Spliting the buffer string Failed")
	}
	s.cpuUsageTimeText = tlst[0]
	return nil
}

func (s *Stat) splitCpuUsageTimeString() ([]string, error) {
	slst := strings.Split(s.cpuUsageTimeText, " ")
	if 12 > len(slst) {
		return []string{}, errors.New("Parse data is invalid data")
	}
	return slst, nil
}

func (s *Stat) getCpuUsageTime(slst []string) error {
	for i := 0; i < len(slst); i++ {
		if regexp.MustCompile(`[0-9]`).Match([]byte(slst[i])) {
			t, err := strconv.Atoi(slst[i])
			if err != nil {
				return err
			}
			s.cpuUsageTimeList = append(s.cpuUsageTimeList, t)
		}
	}

	return nil
}
