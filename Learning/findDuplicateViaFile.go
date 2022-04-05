package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	for _, fileName := range files {
		fh, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while open a file: %v\n", err)
			continue
		}
		countLines(fh, counts)
		fh.Close()
	}

	for index, value := range counts {
		if value > 1 {
			fmt.Printf("Occurance: %d, String: %s\n", value, index)
		}
	}
}

func countLines(fh *os.File, counts map[string]int) {
	input := bufio.NewScanner(fh)
	for input.Scan() {
		counts[input.Text()]++
	}
}
