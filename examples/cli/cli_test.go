package main

import "testing"

func TestCLI_PASS(t *testing.T) {
	if result := Run(); result == 2 {
		t.Log(result)
	} else {
		t.Error("result :", result)
	}
}

func TestCLI_Add(t *testing.T) {
	*Hours = "1:00,00:01,00:02,00:03,00:04,00:05,00:06,00:07,00:08,00:09,00:10,00:11"
	if result := Run(); result == 0 {
		t.Log(result)
	} else {
		t.Error("result :", result)
	}
}

func TestCLI_Subtraction(t *testing.T) {
	*Hours = "1:00,00:01,00:02,00:03,00:04,00:05,00:06,00:07,00:08,00:09,-00:10,-00:11"
	*Subtraction = true
	if result := Run(); result == 0 {
		t.Log(result)
	} else {
		t.Error("result :", result)
	}
}
