package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	hc "github.com/jsbry/go-hourcalc"
)

var (
	Hours       = flag.String("h", "", "hh:mm形式をカンマ区切りで指定")
	Subtraction = flag.Bool("s", false, "-sをつけると減算します")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, helpMessage())
		flag.PrintDefaults()
	}

	os.Exit(Run())
}

func Run() int {
	flag.Parse()

	if *Hours == "" {
		fmt.Println(`00:00`)
		return 2
	}

	hours := strings.Split(*Hours, ",")
	hour := "00:00"
	if *Subtraction == false {
		hour, _ = hc.AddTime(hours)
	} else {
		hour, _ = hc.DiffTime(hours)
	}
	fmt.Println(hour)
	return 0
}

func helpMessage() string {
	return `
Usage of hourcalc:

/* ja */
hh:mm形式をカンマ区切りで渡すと時間を加減します

Options
`

}
