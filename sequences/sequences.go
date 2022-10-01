package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) []int {
	slice = append(slice, slice...)
	return slice
}

func mapSlice(f func(a int) int, slice []int) {
	for i, num := range slice {
		slice[i] = f(num)
	}
}

func mapArray(f func(a int) int, array [5]int) [5]int {
	for i, num := range array {
		array[i] = f(num)
	}
	return array
}

func main() {
	intsSlice := []int{1, 2, 3, 4, 5}
	mapSlice(addOne, intsSlice)
	fmt.Println(intsSlice)

	intsArray := [5]int{1, 2, 3, 4, 5}
	temp := mapArray(addOne, intsArray)
	fmt.Println(temp)

	newSlice := intsSlice[1:3]
	mapSlice(square, newSlice)
	fmt.Println("newSlice: ", newSlice, " intsSlice: ", intsSlice)

	fmt.Println("intsSlice: ", intsSlice)
	intsSlice = double(intsSlice)
	fmt.Println("Doubled intsSlice: ", intsSlice)
}
