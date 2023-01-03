package main

import "fmt"

func main() {
	fmt.Println("Starting the code!")
	greeting_words := greeting("Burak")
	fmt.Println(greeting_words)
}

func greeting(name string) string {
	if len(name) == 0 {
		return "Hello Anonymous!"
	} else {
		return fmt.Sprintf("Hello %s", name)
	}
}
