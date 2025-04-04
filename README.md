# Bloom Filter Analysis in Go

This repository implements Bloom filters in Go and analyzes their performance by varying the number of hash functions and dataset sizes. Bloom filters are space-efficient probabilistic data structures used to test whether an element is possibly in a set or definitely not in a set. This project focuses on understanding how the false positive rate (FPR) changes with different configurations.


## Overview
Bloom filters offer a fast and space-efficient solution for membership queries with the trade-off of having false positives but no false negatives. In this project, I implemented a Bloom filter using multiple hash functions in Go and performed experiments to measure and plot the FPR across varying dataset sizes and hash function counts.

## Features
1. Multiple Hash Functions: Implementation of Bloom filters using different numbers of hash functions.

2. Performance Analysis: Evaluate false positive rates as a function of dataset size and the number of hash functions.

3. Visualization: Plotting capabilities to visualize how performance metrics change with varying configurations.

4. Random Dataset Generation: Uses randomly generated strings to simulate real-world usage.

# How It Works

## Bloom Filter Implementation
BloomFilterMultiHash:
The core data structure that uses a bit array to represent membership.

AddMultiHash & ExistsMultiHash:
Functions to insert keys into the Bloom filter and check for their presence using a specified number of hash functions.

Hash Functions:
Multiple Murmur3-based hash functions are generated with different seeds to provide independent hash outputs.
