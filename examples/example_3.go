package main

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	<-ch
}
