package hourcalc

import "testing"

func TestPASS(t *testing.T) {
	arr := []string{"a", "b", "c"}

	if result, e := AddTime(arr); e != nil {
		t.Error("ERROR : ",e)
	}else{
		t.Log(result)
	}
}

func TestFail(t *testing.T) {
	arr := []string{"a", ""}

	if result, e := AddTime(arr); e != nil {
		t.Error("ERROR : ",e)
	}else{
		t.Log(result)
	}
}

func Test_AddTime_1(t *testing.T) {
	arr := []string{"22:00", "01:00"}

	if result, e := AddTime(arr); e != nil {
		t.Error("ERROR : ",e)
	}else{
		t.Log(result)
	}
}
