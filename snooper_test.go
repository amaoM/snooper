package main

import (
	"bytes"
	"testing"
)

func TestCalculate(t *testing.T) {
	flst := []int{5246691, 72, 7172275, 624873, 1385, 2, 29873, 0, 0, 0}
	slst := []int{5246790, 72, 7172382, 627286, 1385, 2, 29873, 0, 0, 0}
	buf := &bytes.Buffer{}
	ip := "192.168.1.1"

	calculate(flst, slst, buf, ip)

	actual := buf.String()
	expected := `+++++ 192.168.1.1 +++++
user      3 %
nice      0 %
system    4 %
idle     92 %
iowait    0 %
irq       0 %
softirq   0 %
steal     0 %
guest     0 %
+++++++++++++++++++++++++
`

	if actual != expected {
		t.Errorf("expected : %s", expected)
		t.Errorf("actual : %s", actual)
	}
}
