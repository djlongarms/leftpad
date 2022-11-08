package main

import (
	"fmt"
	"github.com/djlongarms/leftpad"
)

func main() {
	fmt.Println((leftpad.Format("hello", 15)))
	fmt.Println((leftpad.Format("goodbye", 15)))
	fmt.Println((leftpad.Format("¿Còmo està?", 15)))
	fmt.Println((leftpad.Format("Internationalization", 15)))
	fmt.Println((leftpad.FormatRune("hello", 15, '😊')))
	fmt.Println((leftpad.FormatRune("goodbye", 15, '😊')))
	fmt.Println((leftpad.FormatRune("¿Còmo està?", 15, '😊')))
}