package main

import "fmt"

func main() {

	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}

	for key, _ := range []int{1, 2, 3, 4} {
		fmt.Printf("key:%d \n", key)
	}
}