// Kanika Warman - mf1513
// CS 311: Assignment 1

package main

import "fmt"

// Function CalcPressure which takes 3 parameter: v,n,t and will returns the pressure p = nRt/v. All are float64.
// ****************** Function CalcPressure begins ******************

func CalcPressure(v, n, t float64) float64 {
	var p float64
	const R = 8.3144598
	fmt.Println("Pressure calculated for below parameters:")
	fmt.Printf("Constant R = %v, Volume = %f, Amount = %f, Temperature = %f\n", R, v, n, t)
	p = (n * R * t) / v
	return p
}

// ****************** Function CalcPressure ends ******************

// Function Decode takes 2 parameters: e and c and uses the code c to return the decrypted version of e
// ****************** Function Decode begins ******************

func Decode(encrypt_text []byte, code map[byte]byte) []byte {
	fmt.Println("The Encrypted text provided is:", string(encrypt_text))
	for i, text := range encrypt_text {
		if value, key := code[text]; key { //check the byte of encrypt_text in the code map as a key and find its value.
			encrypt_text[i] = value
		}
	}
	return encrypt_text
}

// ****************** Function Decode ends ******************

// Function OddParity that takes a slice of ints and returns two values.
//The first return value is true if the list has an odd number of 1’s, false if it has an even number of 1’s.
//The second return value is true if all values are 0’s or 1’s, false if there are any other values.
// ****************** Function OddParity begins ******************

func OddParity(x []int) (bool, bool) {
	fmt.Println("The Slice provided is: ", x)
	result1 := true
	result2 := true
	count_1 := 0 //count of 1's in the array
	for i := 0; i < len(x); i++ {
		if x[i] == 1 {
			count_1++
		} else if x[i] != 0 {
			result2 = false //encountered other than 0/1
		}
	}
	if count_1%2 == 0 {
		result1 = false
	}
	return result1, result2
}

// ****************** Function OddParity ends ******************

func main() {
	fmt.Println(" ***** Calling the first function: CalcPressure *****")
	fmt.Println("Pressure p = ", CalcPressure(1.0, 1.0, 298.15))
	fmt.Println("Pressure p = ", CalcPressure(5.4, 0.36, 98.53))
	fmt.Println("")

	fmt.Println(" ***** Calling the second function: Decode  *****")
	code := map[byte]byte{'e': 'u', 'h': 'f', 'l': 'n', 'o': 'y'}
	fmt.Println("The decrypted text is:", string(Decode([]byte("hello"), code)))
	code = map[byte]byte{'t': 'h', 'h': 'o', 'g': 'l', 'o': 'e'} //no code provided for u!
	fmt.Println("The decrypted text is:", string(Decode([]byte("tough"), code)))
	fmt.Println("")

	fmt.Println(" ***** Calling the third function: OddParity  *****")
	fmt.Println(OddParity([]int{0, 1, 1, 1}))               //true true
	fmt.Println(OddParity([]int{12, 45, 0, 56, 43, 1, 45})) //true false

}
