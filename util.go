package main

import "fmt"

func check[T any](val T, err error) T {
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	return val
}
