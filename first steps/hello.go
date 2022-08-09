// go run hello.go
package main

import "fmt"

func main() {
	fmt.Println("Hello, what is your name?")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Hello, %s.", name)
}
