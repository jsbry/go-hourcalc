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
}
