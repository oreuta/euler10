// getconf.go
package config

import "errors"

const NMAX = 2000000

func Getconf(n int) (int, error) {
	if n == 0 {
		return NMAX, nil
	}
	if n > NMAX {
		return -1, errors.New("number is too big")
	}
	return n, nil
}
