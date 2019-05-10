package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("I have a problem and I think concurrency is the solution")

	sentence := []string{"And", "now", "I", "have", "two", "problems", ".\n"}

	wg.Add(len(sentence))
	for i := 0; i < len(sentence); i++ {
		go printFunc(" " + sentence[i])
	}

	fmt.Println("Now we wait")
	wg.Wait()
}

func printFunc(s string) {
	defer wg.Done()
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
	fmt.Print(s)
}
