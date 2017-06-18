package main

import (
	"bytes"
	"testing"
)

func TestBufferToString(t *testing.T) {
	s := new(Stat)

	s.buffer = *bytes.NewBufferString("cpu  543942 66 704894 622360 1385 2 3770 0 0\n")
	actual, err := s.changeBufferToString()
	expected := "cpu  543942 66 704894 622360 1385 2 3770 0 0\n"
	if err != nil {
		t.Fatal(err)
	}
	if actual != expected {
		t.Errorf("expected : %s", expected)
		t.Errorf("actual : %s", actual)
	}

	s.buffer = *bytes.NewBufferString("")
	_, err = s.changeBufferToString()
	if err == nil {
		t.Error("Got no error")
	}
}
