package main

import "fmt"

func main() {
	var data int

	go func() {
		data++
	}()

	go func() {
		fmt.Println(data)
	}()
}
