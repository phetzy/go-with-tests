package main

import "fmt"

const englishHello = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return englishHello + name
}

func main() {
	fmt.Println(Hello("world"))
}
