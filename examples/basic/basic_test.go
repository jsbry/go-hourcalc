package main

import "testing"

func TestPASS(t *testing.T) {
	hours := []string{"9:45", "10:00", "8:00", "9:00", "", "2:00", "5:00"}

	if result := Run(hours); result == 0 {
		t.Log(result)
	} else {
		t.Error("result :", result)
	}
}
