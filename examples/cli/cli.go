package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	hc "github.com/jsbry/go-hourcalc"
)

var (
	Hours = flag.String("h", "", "RSS Address")
)

func main() {
	os.Exit(Run())
}

func Run() int {
	flag.Parse()

	if *Hours == "" {
		fmt.Println(`command error ===> please hours ' -h "Hours"'`)
		return 2
	}

	hours := strings.Split(*Hours, ",")
	hour, err := hc.AddTime(hours)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hour)
	return 0
}
