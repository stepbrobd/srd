package main

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	go func() {
		ch <- 2
	}()

	<-ch
	<-ch
}
