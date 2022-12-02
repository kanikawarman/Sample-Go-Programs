// Kanika Warman - mf1513
//Assignment 6 - client interface --> a simple command-line user interface that will use the service to make bids on items for sale

package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// port of service to send requests
var PORT = 8080

// Add an item data to the map for the auction
func add(item string, desc string, min_bid_amt int) {
	request := fmt.Sprintf("http://localhost:%v/add?item_name=%v&desc=%v&min_bid=%v", PORT, url.PathEscape(item), url.PathEscape(desc), url.PathEscape(string(min_bid_amt)))
	_, err := http.Get(request)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
}

// Add a bid for the auction item
func bid(scanner *bufio.Scanner, username string) {
	fmt.Print("Name of item to bid on: ")
	scanner.Scan()
	bid_item := scanner.Text()
	fmt.Print("Amount to bid, in $US: ")
	scanner.Scan()
	bid_amount := strings.ToUpper(scanner.Text())
	request := fmt.Sprintf("http://localhost:%v/bid?item_name=%v&bid_amt=%v&user_name=%v", PORT, url.PathEscape(bid_item), url.PathEscape(bid_amount), url.PathEscape(username))
	resp, err := http.Get(request)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	output := string(body)
	if "1" == output[1:2] {
		fmt.Printf("%v is below the minimum bid of $%v\n", bid_amount, output[2:len(output)-1])
	} else if "2" == output[1:2] {
		fmt.Printf("%v is not greater than the current high bid of $%v\n", bid_amount, output[2:len(output)-1])
	} else {
		return
	}
}

// Lookup the item for auction
func lookup(scanner *bufio.Scanner) {
	fmt.Print("Please enter the name: ")
	scanner.Scan()
	item_name := scanner.Text()
	request := fmt.Sprintf("http://localhost:8080/lookup?item_name=%v", url.PathEscape(item_name))
	resp, err := http.Get(request)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if "" == string(body) {
		fmt.Printf("%v is not in the auction list\n", item_name)
	} else {
		fmt.Printf("The item details are: %v\n", string(body))
	}
}

func main() {
	//client := new(http.Client)

	scanner := bufio.NewScanner(os.Stdin)
	var user_name string
	var option string

	fmt.Println("Client for auction service, alpha version\n")
	fmt.Print("Please enter your username: ")
	fmt.Scanf("%s", &user_name)

	quit := "no"
	for quit == "no" {
		fmt.Printf("Bid or Lookup (B/L)?: ")
		fmt.Scanf("%s ", &option)
		if "B" == strings.ToUpper(option) {
			bid(scanner, user_name)
			fmt.Print("\n")
		} else if "L" == strings.ToUpper(option) {
			lookup(scanner)
		}
		fmt.Print("Do you want to quit? (yes/no): ")
		fmt.Scanf("%s ", &quit)
	}

}
