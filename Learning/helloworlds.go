// This is a Hello World program to begin the learning process for Go language

package main

import (
	"os"
	"strings"
	"time"
)

// Import the library, `fmt` is used for format output or scan input. The library needs to be wrapped by double quote
import "fmt"

func helloWorld() {
	// Go natively support Unicode, so it can deal with all language in the world
	fmt.Println("Hello, 世界")
}

// If we run the program with `go build helloworlds.go alpha beta`, the output will be command name + ` alpha beta` * 3,
// and there could be some newline in between.
//
func echo() {
	// Can be replaced with `s := ""` / `var s = "" / var s string = ""`, recommended using the current or the first one
	var s string
	// Can be replaced with `var sep = " "` / `var sep string = " "`, recommended using the current or the first one
	sep := " "
	newline := "\n"

	// Output the command name
	s += os.Args[0]

	start := time.Now()

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		s += newline
	}

	microSecs := time.Since(start).Microseconds()
	fmt.Printf("First takes %d\n", microSecs)

	// start is a new variable, so not use declaration `:=` here
	start = time.Now()
	// The range here will generate two values: index and the value. We don't need index here.
	// Go language does not allow var defined but not used, thus we can not use things like `for temp, arg`
	//
	for _, arg := range os.Args[1:] {
		s += sep + arg
		s += newline
	}

	microSecs = time.Since(start).Microseconds()
	fmt.Printf("Second takes %d\n", microSecs)

	start = time.Now()

	// This will be more efficient compared with first echo loop
	s += sep + strings.Join(os.Args[1:], " ")

	microSecs = time.Since(start).Microseconds()
	fmt.Printf("Third takes %d\n", microSecs)

	fmt.Println(s)
}

func main() {
	helloWorld()
	echo()
}
