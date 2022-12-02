// Kanika Warman - mf1513
//Assignment 6 - service interface --> a simple a simple service to support auctions of client interface

package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

// Describes valid subpaths following
var validPath = regexp.MustCompile(`^/(bid|lookup|add)$`)

// Port for clients to connect to
var PORT_STR = ":8080"

type bidstruct struct {
	itemname string
	desc     string
	min_bid  int
	user     bidder
}

type bidder struct {
	best_bid int
	username string
}

// Takes a function that handles HTTP requests
// Wraps call to this function with code to ensure
// URL is valid
func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}

// auction items are stored in this map
var auctionTable map[string]bidstruct

// adds new auction item to the map
func addHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("user_name")
	item_name := r.URL.Query().Get("item_name")
	desc := r.URL.Query().Get("desc")
	min_bid, _ := strconv.Atoi(r.URL.Query().Get("min_bid"))
	best_bid, _ := strconv.Atoi(r.URL.Query().Get("best_bid"))
	auctionTable[item_name] = bidstruct{itemname: item_name, desc: desc, min_bid: min_bid, user: bidder{best_bid: best_bid, username: username}}
}

// makes bid to the existing items
func bidHandler(w http.ResponseWriter, r *http.Request) {
	result := 0
	item_name := r.URL.Query().Get("item_name")
	user_name := r.URL.Query().Get("user_name")
	bid_amt, _ := strconv.Atoi(r.URL.Query().Get("bid_amt"))
	if item_exists, ok := auctionTable[item_name]; ok {
		amt_best_bid := item_exists.user.best_bid
		amt_min_bid := item_exists.min_bid
		if amt_best_bid > 0 {
			if bid_amt > amt_best_bid {
				auctionTable[item_name] = bidstruct{itemname: item_name, desc: item_exists.desc, min_bid: item_exists.min_bid, user: bidder{best_bid: bid_amt, username: user_name}}
				result = 0
				send := [2]int{result, amt_min_bid}
				fmt.Fprint(w, send)
			} else {
				result = 2
				send := [2]int{result, amt_best_bid}
				fmt.Fprint(w, send)
			}
		} else if bid_amt >= amt_min_bid {
			auctionTable[item_name] = bidstruct{itemname: item_name, desc: item_exists.desc, min_bid: item_exists.min_bid, user: bidder{best_bid: bid_amt, username: user_name}}
			result = 0
			send := [2]int{result, amt_min_bid}
			fmt.Fprint(w, send)
		} else {
			result = 1
			send := [2]int{result, amt_min_bid}
			fmt.Fprint(w, send)
		}
	}
}

// Looksup the auction items based on item name
func lookupHandler(w http.ResponseWriter, r *http.Request) {
	item_name := r.URL.Query().Get("item_name")
	result := auctionTable[item_name]
	fmt.Fprint(w, result)
}

func main() {
	auctionTable = make(map[string]bidstruct)
	http.HandleFunc("/lookup", makeHandler(lookupHandler))
	http.HandleFunc("/bid", makeHandler(bidHandler))
	http.HandleFunc("/add", makeHandler(addHandler))

	log.Fatal(http.ListenAndServe(PORT_STR, nil))
}
