// Kanika Warman - mf1513
//Assignment 7 - service interface --> a simple service to support client payments and account balance checking

package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"sync"
)

// Describes valid subpaths following
var validPath = regexp.MustCompile(`^/(get|set|inc|pay)$`)

// Port for clients to connect to
var PORT_STR = ":8080"

type AccountStore struct {
	sync.Mutex
	owner    string
	balance  int
	counters map[string]AccountStore
}

// get function to get the details of the account holder:  http://localhost:8080/get?name=B
func (cs *AccountStore) get(w http.ResponseWriter, r *http.Request) {
	log.Printf("get %v", r)
	cs.Lock()
	defer cs.Unlock()
	name := r.URL.Query().Get("name")
	if val, ok := cs.counters[name]; ok {
		fmt.Fprintf(w, "%s: owner = %s, balance = $%d\n", name, val.owner, val.balance)
	} else {
		fmt.Fprintf(w, "%s not found\n", name)
	}
}

// set function updates the balance value to the new amount:  http://localhost:8080/set?name=B&balance=9090
func (cs *AccountStore) set(w http.ResponseWriter, req *http.Request) {
	log.Printf("set %v", req)
	cs.Lock()
	defer cs.Unlock()
	name := req.URL.Query().Get("name")
	val := req.URL.Query().Get("balance")
	intval, err := strconv.Atoi(val)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
	} else {
		if user_exists, ok := cs.counters[name]; ok {
			cs.counters[name] = AccountStore{owner: user_exists.owner, balance: intval}
			fmt.Fprintf(w, "ok\n")
			account := cs.counters[name]
			fmt.Fprintf(w, "%s: owner = %s, balance = $%d\n", name, account.owner, account.balance)
		}
	}
}

// inc function increases the balance amount by the mentioned amount. http://localhost:8080/inc?name=B&amt=10
func (cs *AccountStore) inc(w http.ResponseWriter, req *http.Request) {
	log.Printf("inc %v", req)
	cs.Lock()
	defer cs.Unlock()
	name := req.URL.Query().Get("name")
	val, _ := strconv.Atoi(req.URL.Query().Get("amt"))
	if user_exists, ok := cs.counters[name]; ok {
		cs.counters[name] = AccountStore{owner: user_exists.owner, balance: (user_exists.balance + val)}
		fmt.Fprintf(w, "ok\n")
		account := cs.counters[name]
		fmt.Fprintf(w, "%s: owner = %s, balance = $%d\n", name, account.owner, account.balance)
	} else {
		fmt.Fprintf(w, "%s not found\n", name)
	}
}

// pay function transfers the mentioned amount from one user to another. If the user does not have enough money to transfer, then error is generated.
// http://localhost:8080/pay?from=A&to=B&amt=20

func (cs *AccountStore) pay(w http.ResponseWriter, req *http.Request) {
	log.Printf("inc %v", req)
	cs.Lock()
	defer cs.Unlock()
	from := req.URL.Query().Get("from")
	to := req.URL.Query().Get("to")
	val, _ := strconv.Atoi(req.URL.Query().Get("amt"))
	if user_exists, ok := cs.counters[from]; ok {
		if user_exists.balance >= val { //checks if enough balance is present in the account
			if to_user_exists, ok := cs.counters[to]; ok {
				cs.counters[to] = AccountStore{owner: to_user_exists.owner, balance: (to_user_exists.balance + val)}
				cs.counters[from] = AccountStore{owner: user_exists.owner, balance: (user_exists.balance - val)}
				fmt.Fprintf(w, "ok\n")
			}
		} else {
			fmt.Fprintf(w, "%s only has a balance of $%d, cannot pay $%d", from, user_exists.balance, val)
		}
	}
}

func main() {

	//setting up the values in the map for the account holders
	store := AccountStore{counters: map[string]AccountStore{"A": {owner: "Ali", balance: 100}, "B": {owner: "Bob", balance: 50}, "C": {owner: "Chad", balance: 250}}}
	http.HandleFunc("/get", store.get)
	http.HandleFunc("/set", store.set)
	http.HandleFunc("/inc", store.inc)
	http.HandleFunc("/pay", store.pay)

	log.Fatal(http.ListenAndServe(PORT_STR, nil))
}
