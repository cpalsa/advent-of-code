package util

import (
	"cmp"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func LoadInput(fname string) (input []byte, err error) {
	binPath, err := os.Executable()
	if err != nil {
		return input, err
	}

	binDir := filepath.Dir(binPath)
	input, err = os.ReadFile(binDir + "/" + fname)

	return input, err
}

func HandleFatal(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// Returns the minimum value of given Orderables
func Min[T cmp.Ordered](vals ...T) (min T, err error) {
	if len(vals) == 0 {
		return min, errors.New("no comparable values")
	}

	min = vals[0]
	for _, val := range vals {
		if val < min {
			min = val
		}
	}

	return min, nil
}
