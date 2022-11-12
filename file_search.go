//This program does a directory search, listing all line numbers for each file with a specific string using goroutines and channels.

package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// create the go channel for file search
var filesearch = make(chan string, 50)

// function for error output
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// walkpath the directory search recursively
// calls the search file function using gochannel filesearch
func visit(path string, di fs.DirEntry, err error) error {
	go search_file(path, di.Name(), filesearch)
	return nil
}

// called using the go channel for directory search. Uses the WalkDir method to do the directory search recursively
func read_dir(path string, dir_search chan string) {
	error := filepath.WalkDir(path, visit)
	check(error)
}

// Does the file search for the specified string. This is also concurrently done using channel filesearch
func search_file(path string, file_name string, filesearch chan string) {
	file, err := os.Open(path)
	check(err)

	// Splits on newlines by default.
	scanner := bufio.NewScanner(file)
	line := 1 //to diplay the line number in each file
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "here") { //check the string present in the open file
			fmt.Print(file_name, " (", line, "): ")
			fmt.Println(scanner.Text())
		}
		line++
	}
	file.Close()
}

func main() {
	fmt.Println("*****************Starting the File search for string \"here\" in the directory using walkpath recursive directory search!*****************")
	path := "/Users/kanikawarman/Desktop/Fall 2022/CS 311"
	//create the go channel for dir search
	dir_search := make(chan string, 10)
	go read_dir(path, dir_search)

	time.Sleep(5 * time.Second)
	fmt.Println("*****************Ending the File search!*****************")
}
