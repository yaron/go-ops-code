package main

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
)

func main() {
	_, err := git.PlainClone("/tmp/go-ops", false, &git.CloneOptions{
		URL:      "https://github.com/yaron/go-ops.git",
		Progress: os.Stdout,
	})

	if err != nil {
		fmt.Println(err)
	}
}
