package main

import (
	"github.com/spaolacci/murmur3"
)

const (
	numHashFunctions = 100
)

func hashMurmur3MultiHash(key string, hashFnIndex int) uint32 {
	h := murmur3.New32WithSeed(uint32(155 + hashFnIndex*1000))
	h.Write([]byte(key))
	return h.Sum32()
}

type BloomFilterMultiHash struct {
	bitArray []uint8
	size     uint
}

func NewBloomFilterMultiHash(size uint) *BloomFilterMultiHash {
	return &BloomFilterMultiHash{
		bitArray: make([]uint8, (size+7)/8),
		size:     size,
	}
}

func AddMultiHash(key string, bloom *BloomFilterMultiHash, ind int) {

	for i := 0; i < ind; i++ {
		hash := hashMurmur3MultiHash(key, i) // Calculate the hash of the key
		index := hash % uint32(bloom.size)   // Modulo operation to get the index within the Bloom filter size
		ind1 := index / 8                    // Calculate the byte index in the bit array
		ind2 := index % 8                    // Calculate the bit index within the byte
		bloom.bitArray[ind1] |= (1 << ind2)  // Set the bit at the calculated index to 1
	}

}

func ExistsMultiHash(key string, bloom *BloomFilterMultiHash, ind int) bool {

	for i := 0; i < ind; i++ {
		hash := hashMurmur3MultiHash(key, i)
		index := hash % uint32(bloom.size)
		ind1 := index / 8
		ind2 := index % 8

		if (bloom.bitArray[ind1] & (1 << ind2)) == 0 {
			return false
		}
	}

	return true
}
