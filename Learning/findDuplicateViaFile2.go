package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, fileName := range os.Args[1:] {
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when read file: %v\n", err)
			continue
		}
		for _, value := range strings.Split(string(data), "\n") {
			counts[value]++
		}
	}

	for index, value := range counts {
		if value > 1 {
			fmt.Printf("Occurance: %d, String: %s\n", value, index)
		}
	}
}
