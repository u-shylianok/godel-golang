package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	x := 0
	for x < 20 {
		go func() {
			mu.Lock()
			x++
			mu.Unlock()
		}()
		go func() {
			mu.Lock()
			if x%2 == 0 {
				time.Sleep(1 * time.Millisecond)
				fmt.Println(x)
			}
			mu.Unlock()
		}()
	}
}
