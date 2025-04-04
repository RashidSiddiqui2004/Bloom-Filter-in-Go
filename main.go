package main

import (
	"fmt"

	"time"
)

func main() {

	fmt.Println("\t\tImplementing Bloom filter in Go")

	bloomFilterSize := []int{16, 32, 64, 128, 256, 512, 1024, 2048}

	datasetSize := []int{40, 60, 80, 100, 120, 160, 200, 300, 400, 500, 600, 900, 1000}

	trainTestSplit := 0.8

	for _, size := range bloomFilterSize {

		xlist := []int(make([]int, 0))
		ylist := []float64(make([]float64, 0))

		for _, dtsize := range datasetSize {

			dtsize = (size * dtsize) / 100

			xlist = append(xlist, dtsize)

			fmt.Println("Bloom filter size: ", size, "Dataset size: ", dtsize)

			seed := uint32(time.Now().UnixNano())

			bloom := NewBloomFilter(uint(size), seed)

			keys := []string(make([]string, 0))

			for i := 0; i < max(1, int(float64(dtsize)*(trainTestSplit))); i++ {
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

			absentKeys := make([]string, 0, dtsize)

			// Generate test keys not added to the Bloom filter
			for i := 0; i < max(1, int(float64(dtsize)*(1.0-trainTestSplit))); i++ {
				absentKeys = append(absentKeys, GetRandomString(100, 70, 20))
			}

			falsePositives := 0
			trueNegatives := 0

			for _, key := range absentKeys {
				if Exists(key, bloom) {
					falsePositives++
				} else {
					trueNegatives++
				}
			}

			totalNegatives := falsePositives + trueNegatives
			var fpr float64
			if totalNegatives > 0 {
				fpr = float64(falsePositives) / float64(totalNegatives)
			} else {
				fpr = 0.0
			}

			fmt.Println("False positives: ", falsePositives, "True negatives: ", trueNegatives)

			fmt.Printf("False Positive Rate: %.6f\n", fpr)

			// ylist = append(ylist, falsePositives)

			ylist = append(ylist, float64(fpr))
		}

		plotTitle := "Bloom filter of size " + fmt.Sprint(size)
		xlabel := "Dataset size"
		ylabel := "False Positive Rate (FPR)"
		imageLocation := "BloomFilterSize" + fmt.Sprint(size)

		GetPlot(plotTitle, xlabel, ylabel, xlist, ylist, "FPR Plots", imageLocation)
	}

}
