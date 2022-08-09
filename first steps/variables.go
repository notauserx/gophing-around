package main

func main() {
	// x's type is inferred by the compiler.
	x := 5
	var y int = 10
	const z = 11

	println(x + y + z)

	// defining multiple vars
	var (
		a = 1
		b = 2
		c = 3
	)

	println(a + b + c)
}
