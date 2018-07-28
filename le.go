package main

import (
	"fmt"

	"github.com/donnol/module3"
)

func Hello() {
	var le = module3.NewLe("liu le")
	fmt.Println("I am", le.String())
}
