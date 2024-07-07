package main

func main() {
	ch1 := make(chan struct{})
	close(ch1)
	ch1 <- struct{}{}
}
