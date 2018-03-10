package main

import "testing"

func TestSqrt(t *testing.T) {
	_, err := Sqrt(10)

	if err != nil {
		t.Error("this shit errored out")
	}

}
