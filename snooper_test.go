package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func TestOpenFile(t *testing.T) {
	f, err := OpenFile(`./fixtures/stat`)
	if err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(f).String() != "*os.File" {
		t.Error("Cannot get the file object")
	}
}

func TestCannotOpenFile(t *testing.T) {
	f, err := OpenFile(`./fixtures/nofile`)
	if err == nil || f != nil {
		t.Error(err)
	}
}

func TestGetText(t *testing.T) {
	f, err := os.Open(`./fixtures/stat`)
	defer f.Close()
	if err != nil {
		t.Error(err)
	}

	actual, err := GetText(f)

	if err != nil {
		t.Error(err)
	}

	expected := "cpu  543942 66 704894 622360 1385 2 3770 0 0 0"

	if actual != expected {
		t.Errorf("actual: %s, expected: %s", actual, expected)
	}
}

func TestCannotGetText(t *testing.T) {
	f, err := os.Open(`./fixtures/empty`)
	defer f.Close()
	if err != nil {
		t.Error(err)
	}

	actual, err := GetText(f)

	if err == nil || actual != "" {
		t.Error("The text is not empty")
	}
}

func TestGetInvalidText(t *testing.T) {
	f, err := os.Open(`./fixtures/invalid`)
	if err != nil {
		t.Error(err)
	}

	actual, err := GetText(f)

	if err == nil || actual != "" {
		t.Error("The text is invalid")
	}
}

func TestSplitText(t *testing.T) {
	s := "cpu  543942 66 704894 622360 1385 2 3770 0 0 0"
	actual, err := SplitText(s)

	if err != nil {
		t.Error(err)
	}

	expected := []string{"cpu", "", "543942", "66", "704894", "622360", "1385", "2", "3770", "0", "0", "0"}
	if !reflect.DeepEqual(actual, expected) {
		t.Error("The result is different from the expected")
	}
}

func TestSplitShortText(t *testing.T) {
	sl := []string{
		"cpu",
		"cpu  ",
		"cpu  543942",
		"cpu  543942 66",
		"cpu  543942 66 704894",
		"cpu  543942 66 704894 622360",
		"cpu  543942 66 704894 622360 1385 2",
		"cpu  543942 66 704894 622360 1385 2 3770",
		"cpu  543942 66 704894 622360 1385 2 3770 0",
		"cpu  543942 66 704894 622360 1385 2 3770 0 0",
	}

	for _, s := range sl {
		actual, err := SplitText(s)
		if err == nil || actual != nil {
			t.Error("Failed to parse text")
		}
	}
}

func TestParseStat(t *testing.T) {
	slst := []string{"cpu", "", "543942", "66", "704894", "622360", "1385", "2", "3770", "0", "0", "0"}
	actual, err := ParseStat(slst)

	if err != nil {
		t.Error(err)
	}

	expected := []int{543942, 66, 704894, 622360, 1385, 2, 3770, 0, 0, 0}
	if !reflect.DeepEqual(actual, expected) {
		t.Error("The result is different from the expected")
	}
}

func TestCalculate(t *testing.T) {
	flst := []int{5246691, 72, 7172275, 624873, 1385, 2, 29873, 0, 0, 0}
	slst := []int{5246790, 72, 7172382, 627286, 1385, 2, 29873, 0, 0, 0}
	buf := &bytes.Buffer{}

	Calculate(flst, slst, buf)

	actual := buf.String()
	expected := `user      3 %
nice      0 %
system    4 %
idle     92 %
iowait    0 %
irq       0 %
softirq   0 %
steal     0 %
guest     0 %
`

	if actual != expected {
		t.Log(actual)
		t.Log(expected)
		t.Error("The Calculation is incorrect")
	}
}

func TestGetData(t *testing.T) {
	fp := `./fixtures/stat`
	actual := GetData(fp)
	expected := []int{543942, 66, 704894, 622360, 1385, 2, 3770, 0, 0, 0}
	if !reflect.DeepEqual(actual, expected) {
		t.Error("The result is different from the expected")
	}
}
