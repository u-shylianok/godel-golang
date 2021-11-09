package main

import (
	"fmt"
	"time"
)

func main() {
	x := 0

	for i := 0; i < 20; i++ {
		go func() {
			x++
		}()
		go func() {
			y := x
			if y%2 == 0 {
				time.Sleep(1 * time.Millisecond)
				fmt.Println(y)
			}
		}()
	}
}
