package main

import (
	"fmt"
	"utils"
)

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("done")
	dd := []int{1, 5, 3}
	utils.BubbleSort(dd)
	fmt.Println("%v", dd)
}
