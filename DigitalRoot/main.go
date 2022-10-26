package main

import "fmt"

func DigitalRoot(n int) int {
	res := 1

	for n != 0 {
		res *= n % 10
		n = n / 10
	}
	// fmt.Println("ser")
	// fmt.Println(res)

	if res/10 != 0 {
		return DigitalRoot(res)
	} else {
		return res
	}
}

func main() {
	fmt.Println(DigitalRoot(16))
	fmt.Println(DigitalRoot(999999999999))
	fmt.Println(DigitalRoot(942))

}
