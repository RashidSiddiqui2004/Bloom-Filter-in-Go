package main

import (
	"fmt"

	"time"
)

func main() {

	fmt.Println("\t\tImplementing Bloom filter in Go")

	bloomFilterSize := []int{16, 32, 64, 128, 256, 512, 1024, 2048, 4096}

	datasetSize := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	trainTestSplit := 0.8

	for _, size := range bloomFilterSize {

		for _, dtsize := range datasetSize {

			dtsize = size * dtsize / 100

			fmt.Println("Bloom filter size: ", size, "Dataset size: ", dtsize)

			seed := uint32(time.Now().UnixNano())

			bloom := NewBloomFilter(uint(size), seed)

			keys := []string(make([]string, 0))

			for i := 0; i < (int)(float64(dtsize)*trainTestSplit); i++ {
				keys = append(keys, GetRandomString(100, 70, 20))
			}

			for _, key := range keys {
				Add(key, bloom)
			}

			truePositives := 0

			for _, key := range keys {
				if Exists(key, bloom) {
					truePositives++
				}
			}

			println("True positives: ", truePositives)

			// fmt.Println("Updated bloom filter :")
			// fmt.Println(bloom.bitArray)

			absent_keys := []string(make([]string, 0))

			for i := 0; i < (int)(float64(dtsize)*(1.0-trainTestSplit)); i++ {
				absent_keys = append(absent_keys, GetRandomString(100, 70, 20))
			}

			falsePositives := 0

			for _, key := range absent_keys {
				if Exists(key, bloom) {
					falsePositives++
				}
			}

			println("False positives: ", falsePositives)
		}

	}

}
