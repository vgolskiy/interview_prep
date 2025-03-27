package src

// Worker function to read chunks and send them to the channel
import (
	"fmt"
	"os"
	"sync"
)

func ReadChunk(file *os.File, offset int64, size int, wg *sync.WaitGroup, chunks map[int][]byte, mu *sync.Mutex, index int) {
	defer wg.Done()

	buffer := make([]byte, size)
	_, err := file.ReadAt(buffer, offset)
	if err != nil {
		fmt.Printf("\nError reading chunk at offset %d: %v\n", offset, err)
		return
	}

	mu.Lock()
	chunks[index] = buffer
	mu.Unlock()
}
