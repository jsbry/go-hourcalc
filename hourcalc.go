package hourcalc

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// AddTime 時分の加算
func AddTime(hours []string) (string, error) {
	var h, m int
	for _, hour := range hours {
		if hour == "" {
			continue
		}
		hh, mm, err := HourCheck(hour)
		if err != nil {
			return "", err
		}
		h += hh
		m += mm
	}
	if m >= 60 {
		add := m / 60
		h += add
		m = m % 60
	}
	return fmt.Sprintf("%d:%02d", h, m), nil
}

// DiffTime 時分の減算
func DiffTime(hours []string) (string, error) {
	var h, m int
	for _, hour := range hours {
		if hour == "" {
			continue
		}
		hh, mm, err := HourCheck(hour)
		if err != nil {
			return "", err
		}
		h -= hh
		m -= mm
	}
	if math.Abs(float64(m)) >= 60 {
		add := m / 60
		h -= int(math.Abs(float64(add)))
		m = int(m % 60)
		m = int(math.Abs(float64(m)))
	}
	return fmt.Sprintf("%d:%02d", h, m), nil
}

// HourCheck 時分の形式チェック
func HourCheck(hour string) (int, int, error) {
	split := strings.Split(hour, ":")
	if len(split) < 2 {
		return 0, 0, errors.New(hour + " ==> format is not xx:xx")
	}

	hh, err := strconv.Atoi(split[0])
	if err != nil {
		return 0, 0, errors.New(split[0] + " ==> HH is not numeric")
	}
	mm, err := strconv.Atoi(split[1])
	if err != nil {
		return 0, 0, errors.New(split[1] + " ==> MM is not numeric")
	}

	return hh, mm, nil
}
