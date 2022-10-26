package main

import "fmt"

func Persistence(n int) int {

	if n/10 == 0 {
		return 0
	}

	counter := 0

	for n/10 != 0 {
		res := 1
		for n/10 != 0 {
			res *= (n % 10)
			n /= 10
			if n/10 == 0 {
				res *= (n % 10)
			}

		}
		n = res
		counter++

		continue
	}

	return counter
}

func main() {
	fmt.Println("res: ")

	fmt.Print(Persistence(25))
}
