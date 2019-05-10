package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fileLocations := make(chan string, 2)

	go fetchFile("http://www.example.com/file1", fileLocations)
	go fetchFile("http://www.example.com/file2", fileLocations)

	file1 := <-fileLocations
	file2 := <-fileLocations
	fmt.Println(file1)
	fmt.Println(file2)
}

func fetchFile(url string, location chan<- string) {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
	urlSplit := strings.Split(url, "/")
	location <- "/tmp/filedir/" + urlSplit[len(urlSplit)-1]
}
