package main

import "testing"

func TestCLI(t *testing.T) {
	if result := Run(); result == 2 {
		t.Log(result)
	} else {
		t.Error("result :", result)
	}
}
