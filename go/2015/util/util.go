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

func PrintSlice[T any](s []T) {
	// Not sure how to make this generic to only range-able to also handle maps
	fmt.Println("slice[")
	for k, v := range s {
		fmt.Printf("  %v: %+v,\n", k, v)
	}
	fmt.Println("]")
}

func PrintMap[K comparable, V any](m map[K]V) {
	// Not sure how to make this generic to only range-able to also handle slices
	fmt.Println("map{")
	for k, v := range m {
		fmt.Printf("  %v: %+v,\n", k, v)
	}
	fmt.Println("}")
}

// Returns the minimum value of given Orderables
// DEPRECATED -- go 1.21 also introduced min/max funcs
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
