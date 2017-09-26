package hourcalc

import (
	"errors"
)

func AddTime(hours []string) (string, error) {
	for _, hour := range hours {
		if hour == "" {
			return "false", errors.New("空文字です")
		}
	}
	return "true", nil
}
