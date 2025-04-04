package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Observation

// As we increase the number of hash functions, FPR increases
// Optimal value of number of hash functions is given by the formula -
// k = (m/n) * ln2

// 𝑛 = Number of inserted elements

// 𝑚 = Number of bits in the Bloom filter

// 𝑘 = Number of hash functions

func main() {

	fmt.Println("\t\tImplementing Bloom filter in Go")

	// Initialize the global hash functions
	// Init()

	datasetSizes := []int{10, 20, 30, 40, 50, 60, 70}
	trainTestSplit := 0.4
	bloomFilterSize := 1024

	// Loop over different numbers of hash functions (from 1 to 10)
	for numHashes := 1; numHashes <= 30; numHashes++ {

		var xlist []int
		var ylist []float64

		for _, pct := range datasetSizes {

			currentDatasetSize := (bloomFilterSize * pct) / 100
			xlist = append(xlist, currentDatasetSize)

			fmt.Println("Bloom filter size:", bloomFilterSize, "Dataset size:", currentDatasetSize)

			bloom := NewBloomFilterMultiHash(uint(bloomFilterSize))

			// Generate training keys that will be added to the filter
			numTrainKeys := max(1, int(float64(currentDatasetSize)*trainTestSplit))
			var trainingKeys []string
			for range numTrainKeys {
				trainingKeys = append(trainingKeys, GetRandomString(100, 70, 20))
			}

			// Insert training keys using the current number of hash functions
			for _, key := range trainingKeys {
				AddMultiHash(key, bloom, numHashes)
			}

			truePositives := 0
			for _, key := range trainingKeys {
				if ExistsMultiHash(key, bloom, numHashes) {
					truePositives++
				}
			}

			numTestKeys := max(1, int(float64(currentDatasetSize)*(1.0-trainTestSplit)))
			var absentKeys []string
			for range numTestKeys {
				absentKeys = append(absentKeys, GetRandomString(100, 70, 20))
			}

			falsePositives := 0
			trueNegatives := 0
			for _, key := range absentKeys {
				if ExistsMultiHash(key, bloom, numHashes) {
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

			fmt.Printf("Hash Functions: %d, FPR: %.6f\n", numHashes, fpr)
			ylist = append(ylist, fpr)
		}

		plotTitle := fmt.Sprintf("Bloom Filter FPR with %d Hash Functions", numHashes)
		xlabel := "Dataset size"
		ylabel := "False Positive Rate (FPR)"
		imageLocation := fmt.Sprintf("HashFunctionsCnt_%d", numHashes)

		GetPlot(plotTitle, xlabel, ylabel, xlist, ylist, "Plots for MultipleHashFunctions/FPR Plots", imageLocation)
	}
}
