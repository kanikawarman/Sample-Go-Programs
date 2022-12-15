# Sample-Go-Programs
This repository contains some basic Golang programs that have been developed as a part of the school project.
Below are the files that are present in this repository:
1) various_functions.go >> This file has some basic go programming functioning. Function 1 works by taking 3 input parameters and returns an output. The function 2 displays the use of arrays and maps working in Go.
2) regex_expressions.go >> This function displays an example of grammar rules.
3) guess_num.go >> This program is a simple working of structs and methods/interfaces. It is a game that lets the user to choose the random guess number or it wants the computer to play. Depending upon the user choice, the respect structs are activated and random_guess number is played.
4) file_search.go >> This function displays the directory and file search approach. It uses the walkDir method to recursively visit all the folders in the mentioned path and then search the files present in the folder with the mentioned search string. If the file is found with the search string, it then displays the file name along with the line number and displays the entire line. This program makes use of Go-Channel to create a new thread to visit the directories and to search the files. 
5) puzzle.go >> Ask the user a puzzle and respond if it is correct or not. Use of interface and struct
6) 6) pairless.go >> Pair up the elements and return a slice of bool -- true if the value from the first slice is less than the value from the second slice, false otherwise.
     The function takes two slices of int’s as parameters.
    o	Pair up the elements and return a slice of bool -- true if the value from the first slice is less than the value from the second slice, false     otherwise
    o	For example:
    slice1: = []int {0, 255, 97, 98}
    slice2: = []int [12, 300, 80, 100]
    pairLess (slice1, slice2) should return [true, true, false, true]
7) client_interface and service_interface together display the basic functioning of the http service provided by Go. The client_interface runs on the user machine and the service_interface runs on the server. For now, no DB has been added to this code. Hence whenever a new auction request is made, it is done by submitting the http request in the browser directly. Once the service_interface is running, the item is registered on the server by submitting a request in the browser, the client_interface is run on the machine.
Description: A simple auction system, where an item added to the auction list via server request and the user can bid on the items or check the items present in the auction list.
8) Account_pay.go>> This shows the use of locks and mutex in Golang. A simple pay service system is created to support client payments and account balance checking, users are already registered using Structs and maps are used to key a person. A person can check the balance, pay to another person, update the balance records.
