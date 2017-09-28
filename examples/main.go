package main

import (
	"fmt"

	hc "github.com/jsbry/go-hourcalc"
)

func main() {
	_, _, err := hc.HourCheck("abc")
	if err != nil {
		fmt.Println(err)
	}

	hours := []string{"9:45", "10:00", "8:00", "9:00", "", "2:00", "5:00"}
	hour, err := hc.AddTime(hours)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hour)
}
