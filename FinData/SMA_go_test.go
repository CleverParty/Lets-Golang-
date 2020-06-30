package main

import "testing"

func TestMain(t *testing.T) {
	total := fact(2)
	if total != 10 {
		t.Errorf("Testing SMA_GO_TEST package")
	}
}
