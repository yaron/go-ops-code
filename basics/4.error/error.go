package main

import "fmt"

type myError struct{ Message string }

func (e myError) Error() string {
	return e.Message
}

func errorFunc() (string, error) {
	return "", myError{Message: "Sorry..."}
}

func main() {
	value, err := errorFunc()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(value)
}
