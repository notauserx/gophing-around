package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	y := [5]int{1, 2, 3, 4, 5}
	fmt.Println(y)

	total := 0
	for _, value := range y {
		total += value
	}

	fmt.Println("total for y -> ", total)

	fmt.Println("Calling printAverage...")
	printAverage()

}

func printAverage() {
	var x [5]float64
	x[0] = 2.0
	x[1] = 1.0
	x[2] = 3.0
	x[3] = 4.0
	x[4] = 111.0

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))
}
