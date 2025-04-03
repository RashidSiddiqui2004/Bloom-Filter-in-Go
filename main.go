package main

import (
	"fmt"

	"time"

	"github.com/spaolacci/murmur3"
)

// hashMurmur3 hashes a string using MurmurHash3 with a given seed.

func hashMurmur3(data string, seed uint32) uint32 {
	h := murmur3.New32WithSeed(seed)
	h.Write([]byte(data))
	return h.Sum32()
}

type BloomFilter struct {
	bitArray []bool
	size     uint
	seed     uint32
}

func NewBloomFilter(size uint, seed uint32) *BloomFilter {
	return &BloomFilter{
		bitArray: make([]bool, size),
		size:     size,
		seed:     seed,
	}
}

func Add(key string, bloom *BloomFilter) {
	hash := hashMurmur3(key, bloom.seed)
	index := hash % uint32(bloom.size)
	fmt.Println(key, " added at index ", index)
	bloom.bitArray[index] = true
}

func Exists(key string, bloom *BloomFilter) bool {
	hash := hashMurmur3(key, bloom.seed)
	index := hash % uint32(bloom.size)
	return bloom.bitArray[index]
}

func main() {
	fmt.Println("Implementing Bloom filter!")
	seed := uint32(time.Now().UnixNano())

	bloom := NewBloomFilter(16, seed)

	keys := []string{"a", "b", "c"}

	for _, key := range keys {
		Add(key, bloom)
	}

	for _, key := range keys {
		fmt.Println(Exists(key, bloom))
	}

	fmt.Println("Updated bloom filter :")

	fmt.Println(bloom.bitArray)

}
