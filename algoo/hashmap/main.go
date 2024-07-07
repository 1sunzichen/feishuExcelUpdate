package main

import "fmt"

type pair struct {
	key int
	val string
}
type arrayHashMap struct {
	buckets []*pair
}

func newArrayHashMap() *arrayHashMap {
	buckets := make([]*pair, 100)
	return &arrayHashMap{buckets: buckets}
}

func (a *arrayHashMap) hashFunc(key int) int {
	index := key % 100
	return index
}

func main() {
	key := "kkkk81212k"
	fmt.Println([]byte(key))
}
func addHash(key string) int {
	var hash int64
	// var modules int64
	modules := int64(1000000007)
	for _, v := range []byte(key) {
		hash = (hash + int64(v)) % modules
	}
	return int(hash)
}
