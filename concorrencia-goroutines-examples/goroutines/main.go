package main

import (
	"fmt"
	"time"
)

func printText(text string) {
    for {
        fmt.Printf("Writing: %s\n", text)
        time.Sleep(time.Second)
    }
}

func main() {
    go printText("Hello World!")
    printText("Akira")
}