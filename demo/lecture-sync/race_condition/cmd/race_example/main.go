package main

import (
	"fmt"
	"time"
)

func main() {
	x := 0
	for x < 20 {
		go func() {
			x++
		}()
		go func() {
			if x%2 == 0 {
				time.Sleep(1 * time.Millisecond)
				fmt.Println(x)
			}
		}()
	}
}
