package main

func InAscOrder(numbers []int) bool {
	result := false
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > numbers[i-1] {
			result = true
		} else {
			result = false
			break
		}
	}

	return result

}

func main() {

	InAscOrder([]int{1, 2, 4, 7, 19})

}
