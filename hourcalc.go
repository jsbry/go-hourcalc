package hourcalc

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func AddTime(hours []string) (string, error) {
	for _, hour := range hours {
		hh, mm, err := HourCheck(hour)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(hh, mm)
	}
	return "10:00", nil
}

func HourCheck(hour string) (int, int, error) {
	split := strings.Split(hour, ":")
	if len(split) < 2 {
		return 0, 0, errors.New("format is not xx:xx")
	}

	hh, err := strconv.Atoi(split[0])
	if err != nil {
		return 0, 0, errors.New("HH is not numeric")
	}
	mm, err := strconv.Atoi(split[1])
	if err != nil {
		return 0, 0, errors.New("MM is not numeric")
	}

	return hh, mm, nil
}
