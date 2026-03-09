package main

import (
	"fmt"
	"github.com/denivladislav/hexlet-go/greeting"
	"github.com/fatih/color"
)

func main() {
	c := color.New(color.FgCyan).Add(color.Underline)
	fmt.Println("The following text is colored:")
	_, err := c.Println(greeting.Hello())
	if err != nil {
		fmt.Println("Error in println!", err)
	}
}
