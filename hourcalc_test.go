package hourcalc

import "testing"

func TestPASS(t *testing.T) {
	arr := []string{"0:0"}

	if result, e := AddTime(arr); e == nil {
		t.Log(result)
	} else {
		t.Error(result, e)
	}
}

func TestFail(t *testing.T) {
	arr := []string{"a", ""}

	if result, e := AddTime(arr); e != nil {
		t.Log(result)
	} else {
		t.Error(result, e)
	}
}

func Test_AddTime_1(t *testing.T) {
	arr := []string{"22:50", "01:50", "02:00:59"}

	if result, e := AddTime(arr); e == nil {
		t.Log(result)
	} else {
		t.Error(result, e)
	}
}
func Test_DiffTime_1(t *testing.T) {
	arr := []string{"22:50", "01:50", "02:00:59"}

	if result, e := DiffTime(arr); e == nil {
		t.Log(result)
	} else {
		t.Error(result, e)
	}
}
