package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronus(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Zakir Ganteng")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronus(group)
	}

	group.Wait()
	fmt.Println("Selesai")
}
