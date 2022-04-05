package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fileMap := make(map[string]map[string]int)
	for _, fileName := range os.Args[1:] {
		counts := make(map[string]int)
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when read file: %v\n", err)
			continue
		}
		for _, value := range strings.Split(string(data), "\n") {
			counts[value]++
		}
		fileMap[fileName] = counts
	}
	var res []string
	for fileName, dupCount := range fileMap {
		fmt.Printf("Going through %s\n", fileName)
		var hasDup bool
		for index, value := range dupCount {
			if value > 1 {
				hasDup = true
				fmt.Printf("Occurance: %d, String: %s\n", value, index)
			}
		}
		if hasDup {
			res = append(res, fileName)
		}
		fmt.Println("--------------------------")
	}
	fmt.Printf("The file that has duplicate are: %v\n", res)
}
