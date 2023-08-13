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

// Iterates over the input slice, passing the index and value to the callback function for each iteration.
// Returns a new slice containing only values that pass the callback predicate (return true)
func FilterSlice[T any](s []T, cb func(i int, v T) bool) []T {
	// Not sure how to make this more generic to also handle maps
	filtered := []T{}
	for i, v := range s {
		if cb(i, v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

// Iterates over the input map, passing the key and value to the callback function for each iteration.
// Returns a new map containing only key+value pairs that pass the callback predicate (return true)
func FilterMap[K comparable, V any](m map[K]V, cb func(k K, v V) bool) map[K]V {
	// Not sure how to make this more generic to also handle slices
	filtered := make(map[K]V)

	for k, v := range m {
		if cb(k, v) {
			filtered[k] = v
		}
	}

	return filtered
}

// Iterates over the input array, passing the index and value to the callback function for each iteration.
// Returns a new array built from the return values of the callback function.
func Map[T any, M any](s []T, cb func(i int, v T) M) []M {
	mapped := []M{}

	for i, v := range s {
		mapped = append(mapped, cb(i, v))
	}

	return mapped
}

func PrintSlice[T any](s []T) {
	// Not sure how to make this generic to only range-able to also handle maps
	fmt.Println("[")
	for k, v := range s {
		fmt.Printf("  %v: %+v,\n", k, v)
	}
	fmt.Println("]")
}

func PrintMap[K comparable, V any](m map[K]V) {
	// Not sure how to make this generic to only range-able to also handle slices
	fmt.Println("{")
	for k, v := range m {
		fmt.Printf("  %v: %+v,\n", k, v)
	}
	fmt.Println("}")
}

// Iterates over a slice to see if the given value exists within the slice.
func SliceContains[T comparable](s []T, v T) bool {
	for _, sv := range s {
		if sv == v {
			return true
		}
	}

	return false
}

// Iterates over a map to see if the given value exists within the map.
func MapContains[K comparable, V comparable](m map[K]V, v V) bool {
	for _, mv := range m {
		if mv == v {
			return true
		}
	}

	return false
}

// Returns a slice containing all keys of the given map
func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := []K{}

	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

// Returns a slice containing all values of the given map
func MapVals[K comparable, V any](m map[K]V) []V {
	vals := []V{}

	for _, v := range m {
		vals = append(vals, v)
	}

	return vals
}

// Returns the minimum value respective to the type of the values given
func Min[T cmp.Ordered](vals ...T) T {
	if len(vals) == 0 {
		HandleFatal(errors.New("Min: no values to compare"))
	}

	if len(vals) == 1 {
		return vals[0]
	}

	min := vals[0]
	for _, v := range vals {
		if v < min {
			min = v
		}
	}

	return min
}

// Returns the maximum value respective to the type of the values given
func Max[T cmp.Ordered](vals ...T) T {
	if len(vals) == 0 {
		HandleFatal(errors.New("Max: no values to compare"))
	}

	if len(vals) == 1 {
		return vals[0]
	}

	max := vals[0]
	for _, v := range vals {
		if v > max {
			max = v
		}
	}

	return max
}
