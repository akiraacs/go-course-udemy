package main

import (
	"fmt"
	"time"
)

func printText(text string, ch chan string) {
	for i := 0; i < 5; i++ {
		ch <- text
		time.Sleep(time.Second)
	}

	close(ch)
}

func main() {
	channel := make(chan string)

	go printText("Hello World!", channel)

	println("Passed here ...")

	// for {
	// 	message, open := <-channel
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(message)
	// }

	for message := range channel {
		fmt.Println(message)
	}

	fmt.Println("Exiting the program ...")

}
