package main

import "fmt"

type myConfig struct {
	APIKey string
	Secret string
}

func main() {
	config := myConfig{
		APIKey: "wo9wmfojw4rfwsmef",
		Secret: "so9jfsofjsgsg",
	}

	fmt.Println(config.APIKey)
}
