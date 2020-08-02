package main

import (
	"fmt"

	"leftpad"
)

func main() {
	fmt.Println(leftpad.Format("hello", 15))
	fmt.Println(leftpad.FormatRune("hello", 15, '😊'))
}
