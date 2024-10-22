package main

import (
	"fmt"
	"sync"
	"time"
)

func printText(text string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Writing: %s\n", text)
		time.Sleep(time.Second)
	}
}

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(5)

	go func() {
		printText("GoRoutine 1")
		waitGroup.Done() // -1
	}()

	go func() {
		printText("GoRoutine 2")
        waitGroup.Done() // -1
	}()

	go func() {
		printText("GoRoutine 3")
        waitGroup.Done() // -1
	}()
    
	go func() {
		printText("GoRoutine 4")
        waitGroup.Done() // -1
	}()
    
	go func() {
		printText("GoRoutine 5")
        waitGroup.Done() // -1
	}()

    waitGroup.Wait()
}
