package main

import "fmt"

func main() {

	var a int = 1
	var b int = 2

	fmt.Println("int demo")
	fmt.Println("a + b = ", a+b)

	fmt.Println("floating point")
	fmt.Println("1.0 + 1.0 =", 1.0+1.0)

	fmt.Println("Strings")
	fmt.Println(len("Hello, World"))
	fmt.Println("Hello, World"[1])
	fmt.Println("Hello, " + "World")

	var t bool = true
	var f bool = false
	fmt.Println("boolean")
	fmt.Println(t && f)
	fmt.Println(t || f)
}
