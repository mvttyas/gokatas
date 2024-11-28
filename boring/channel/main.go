// A channel allows for communication and synchronization between goroutines.
//
// Level: beginner
// Topics: goroutines, channels
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// c := make(chan string)
	// go say("blah", c)
	// generator "pattern" for main x goroutine sync thanks to channel use
	c := say("blop")
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}

}

func say(msg string) <-chan string { //returns recevei-only channel of strings
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s, %d", msg, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}()
	return c
}
