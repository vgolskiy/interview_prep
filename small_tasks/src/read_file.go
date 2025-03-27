package src

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func ReadFileByChunks() {
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)

	chunk := make([]byte, 1)
	for {
		rows, err := f.Read(chunk)
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("\nFinished reading")
			} else {
				fmt.Println(err)
			}
			break
		}
		fmt.Print(string(chunk[:rows]))
	}
	return
}
