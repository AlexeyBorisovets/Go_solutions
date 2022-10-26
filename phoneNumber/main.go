package main

import (
	"strconv"
)

// returns "(123) 456-7890"
func CreatePhoneNumber(numbers [10]uint) string {
	res := "("

	for i := 0; i < 3; i++ {
		res += strconv.Itoa(int(numbers[i]))
	}

	res += ") "

	for i := 3; i < 6; i++ {
		res += strconv.Itoa(int(numbers[i]))
	}

	res += "-"

	for i := 6; i < 10; i++ {
		res += strconv.Itoa(int(numbers[i]))
	}

	return res
}

func main() {
	CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})

}
