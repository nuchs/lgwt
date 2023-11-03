package main

import "fmt"

const englishGreetingPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishGreetingPrefix + name + "!"
}

func main() {
	fmt.Println(Hello("World"))
}
