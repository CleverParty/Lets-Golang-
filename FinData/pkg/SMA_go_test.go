package pkg

import (
	"testing"
)

func fact(num int) int {
	if num > 1 {
		return num * fact(num-1) // fact test fails but main works
	} else {
		return 1
	}
}

func TestfactFunc(t *testing.T) {
	value := fact(3)
	if value != -3 {
		t.Errorf("func fact is throwing %d , check out flow again! ", value)
	}
}
