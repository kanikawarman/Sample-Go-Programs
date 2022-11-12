
//Ask the user a puzzle and respond if it is correct or not

package main

import (
	"fmt"
	"strings"
)

// interface puzzle having solve() function with bool return value
type puzzle interface {
	solve() bool
}

// the struct has question and the correct answer
type riddle struct {
	question string
	answer   string
}

/*solve method takes the struct input and return the bool.
It takes user input and compares the correct asnwer with it, if it is correct, then returns true or else false */

func (r riddle) solve() (response bool) {

	var user_response string
	fmt.Println("Your Question is: ")
	fmt.Println(r.question)
	fmt.Print("Enter your response: ")
	fmt.Scanf("%s", &user_response)
	if r.answer == strings.ToUpper(user_response) { // It turns the user response to uppercase to ignore case sensitivity
		return true
	} else {
		return false
	}
}

func main() {
	var puzz puzzle
	rid := riddle{question: "Which is the coolest alphabet?", answer: "B"}
	puzz = rid
	fmt.Println("Welcome to Riddles!\n")
	result := puzz.solve()
	if result == true {
		fmt.Println("Solved!")
	} else {
		fmt.Println("Wrong answer!")
	}
}
