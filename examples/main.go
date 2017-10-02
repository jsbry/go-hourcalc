package main

import (
	"fmt"
	"os"

	hc "github.com/jsbry/go-hourcalc"
)

func main() {
	arr := []string{"0:0"}
	os.Exit(Run(arr))
}

func Run(hours []string) int {
	hour, err := hc.AddTime(hours)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hour)

	return 0
}
