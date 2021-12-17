package domain

import "golang.org/x/xerrors"

func Plus(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

func Multiple(a, b int) int {
	return a * b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, xerrors.New("invalid value 0")
	}
	return a / b, nil
}
