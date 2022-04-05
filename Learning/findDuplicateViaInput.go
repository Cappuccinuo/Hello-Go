package main

import (
	"bufio"
	"fmt"
	"os"
)

// The find duplicate application will find the line with occurrence larger than 1.
func main() {
	// Create a map to store key/value pair, key is string and value is int
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	fmt.Println("Please entry the strings.")
	// In Go Land Application, there is no way to terminate in the console.
	// When executed in Terminal, the Scan can be ended with ctrl + d
	//
	for input.Scan() {
		counts[input.Text()]++
	}

	for index, value := range counts {
		if value > 1 {
			fmt.Printf("Occurance: %d String: %s\n", value, index)
		}
	}
}
