package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case msg := <-input1:
				c <- msg
			case msg := <-input2:
				c <- msg
			}
		}
	}()
	return c
}

func main() {
	c := fanIn(boring("Jack"), boring("Jill"))
	for i := 0; i < 30; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Exiting")
}
