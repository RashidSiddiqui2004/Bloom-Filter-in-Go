package main

import (
	"github.com/spaolacci/murmur3"
)

// hashMurmur3 hashes a string using MurmurHash3 with a given seed.

func hashMurmur3(data string, seed uint32) uint32 {
	h := murmur3.New32WithSeed(seed)
	h.Write([]byte(data))
	return h.Sum32()
}

type BloomFilter struct {
	bitArray []uint8
	size     uint
	seed     uint32
}

func NewBloomFilter(size uint, seed uint32) *BloomFilter {
	return &BloomFilter{
		bitArray: make([]uint8, (size+7)/8),
		size:     size,
		seed:     seed,
	}
}

func Add(key string, bloom *BloomFilter) {
	hash := hashMurmur3(key, bloom.seed) // Calculate the hash of the key
	index := hash % uint32(bloom.size)   // Modulo operation to get the index within the Bloom filter size
	ind1 := index / 8                    // Calculate the byte index in the bit array
	ind2 := index % 8                    // Calculate the bit index within the byte
	bloom.bitArray[ind1] |= (1 << ind2)  // Set the bit at the calculated index to 1
}

func Exists(key string, bloom *BloomFilter) bool {
	hash := hashMurmur3(key, bloom.seed)
	index := hash % uint32(bloom.size)
	ind1 := index / 8
	ind2 := index % 8
	return (bloom.bitArray[ind1] & (1 << ind2)) != 0
}
