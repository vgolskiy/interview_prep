package main

import (
	"InterviewPrep/assignments/src"
	"fmt"
	"os"
	"sync"
	"time"
)

const chunkSize = 1 // Size of each chunk in bytes

func main() {
	// Open the file
	file, err := os.Open("/Users/Spotlas/Desktop/InterviewPrep/assignments/file_by_chunks_parallel/test.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	// Get file size
	info, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}
	fileSize := info.Size()

	// Use a map with a mutex to store chunks in order
	chunks := make(map[int][]byte)
	var mu sync.Mutex
	var wg sync.WaitGroup

	// Launch goroutines in parallel
	start := time.Now()
	chunkCount := 0
	for offset := int64(0); offset < fileSize; offset += chunkSize {
		wg.Add(1)
		go src.ReadChunk(file, offset, chunkSize, &wg, chunks, &mu, chunkCount)
		chunkCount++
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Print the collected chunks in order
	for i := 0; i < chunkCount; i++ {
		fmt.Print(string(chunks[i]))
	}

	fmt.Printf("\nFile read in parallel in %v\n", time.Since(start))
}
