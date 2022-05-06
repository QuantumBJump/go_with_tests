package main

import "fmt"

func Hello(name string) string {
	var target = ""
	if name != "" {
		target = ", " + name
	}
	return "Hello" + target
}

func main() {
	fmt.Println(Hello("world"))
}
