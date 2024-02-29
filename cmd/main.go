package main

import "fmt"

var version = "dev"

func main() {
	fmt.Printf("Version: %s\n", version)

	fmt.Println(printApplicationStarted())
}

func printApplicationStarted() string {
	return "Growing Buddy application started!"
}
