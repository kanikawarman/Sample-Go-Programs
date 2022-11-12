//Kanika Warman - mf1513
//Pair up the elements and return a slice of bool -- true if the value from the first slice is less than the value from the second slice, false otherwise

package main

import "fmt"

// function takes the slices and comapares it values and returns the bool array
func pairLess(slice1 []int, slice2 []int) []bool {
	fmt.Println("Slice 1 is: ", slice1)
	fmt.Println("Slice 2 is: ", slice2)
	var length = len(slice1)
	bool_array := make([]bool, length) //make a array that will have result in bool
	for i := 0; i < length; i++ {
		if slice1[i] < slice2[i] {
			bool_array[i] = true
		} else {
			bool_array[i] = false
		}
	}
	return bool_array
}

func main() {
	fmt.Println("**************** Example 1 ****************")
	slice1 := []int{0, 255, 97, 98}
	slice2 := []int{12, 300, 80, 100}
	pair_result := pairLess(slice1, slice2)
	fmt.Println("Pair Result: ", pair_result)
	fmt.Println("**************** Example 2 ****************")
	slice3 := []int{10, 25, 970, 198}
	slice4 := []int{1, 300, 80, 100}
	pair_result = pairLess(slice3, slice4)
	fmt.Println("Pair Result: ", pair_result)
}
