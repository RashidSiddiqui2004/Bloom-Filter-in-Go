package main

import (
	"fmt"

	"github.com/spaolacci/murmur3"
)

// hashMurmur3 hashes a string using MurmurHash3 with a given seed.
func hashMurmur3(data string, seed uint32) uint32 {
	h := murmur3.New32WithSeed(seed)
	h.Write([]byte(data))
	return h.Sum32()
}

func main() {
	fmt.Println("Implementing Bloom filter!")

	input := "hello"
	result := hashMurmur3(input, 0)

	fmt.Printf("Hash of '%s': %d\n", input, result)
}
