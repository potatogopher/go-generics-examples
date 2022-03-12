package main

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func MapValues[K comparable, V any](m map[K]V) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

func Print[T any](s ...T) {
	for _, v := range s {
		fmt.Print(v)
	}
	fmt.Println()
}

func Sum[T int64 | int | float64](a, b T) T {
	return a + b
}
