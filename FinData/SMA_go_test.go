package main

import "testing"

func TestMain(t *testing.T) {
	total := testing.Main(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
