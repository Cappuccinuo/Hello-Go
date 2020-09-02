package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    for index, arg := range os.Args {
         fmt.Println(strconv.Itoa(index) + " " + arg)
    }
}
