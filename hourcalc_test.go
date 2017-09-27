package hourcalc

import "testing"

func TestPASS(t *testing.T) {
	arr := []string{"0:0"}

	if result, e := AddTime(arr); e == nil {
		t.Log(result)
	} else {
		t.Error("ERROR : ", e)
	}
}

func TestFail(t *testing.T) {
	arr := []string{"a", ""}

	if result, e := AddTime(arr); e != nil {
		t.Log("ok fail", result)
	} else {
		t.Error("ERROR : ", result, e)
	}
}

func Test_AddTime_1(t *testing.T) {
	arr := []string{"22:50", "01:50", "02:00:59"}

	if result, e := AddTime(arr); e == nil {
		t.Log("ok : ", result)
	} else {
		t.Error("ERROR : ", result, e)
	}
}
