package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 10; i++ {
		switch i {
		case 1:
			fmt.Println("one")
		case 5:
			fmt.Println("five")
		default:
			fmt.Println("something other than one and five")
		}
		sum += i
	}

	fmt.Println(sum)

	i := 100
	sum = 0
	for i <= 200 {
		if i%4 == 0 {
			sum += i
		} else if i%5 == 0 {
			sum = sum * 2
		}
		i += 5
	}

	fmt.Println(sum)
}
