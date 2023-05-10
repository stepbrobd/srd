package main

import "fmt"

func main() {
	var data int
	done := make(chan bool)

	go func() {
		data++
		done <- true
	}()

	go func() {
		<-done
		fmt.Println(data)
	}()
}
