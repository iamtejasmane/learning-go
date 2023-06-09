package main

import "testing"

var tests = []struct {
	name string
	dividend float32
	divisor float32
	expected float32
	isErr bool 
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 10.0, true}
}

func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("Got an error when we should not have")
	}
}
func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err == nil {
		t.Error("Did not get an error when we should have")
	}
}
