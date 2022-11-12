package main

import (
	"fmt"
	"regexp"
)

func main() {
	matched, err := regexp.MatchString(`^([a-zA-Z]+(-[a-zA-Z]+)+) ([a-zA-Z]+(-[a-zA-Z]+)+)$`, "ju-NIP-er-us skop-u-LO-rum")
	fmt.Println(matched, err)
	matched, err = regexp.MatchString(`^(#\d - [1-9]*.\d{2} Gallon \$[1-9][0-9][0-9]*.\d{2})$`, "#1 – .75 Gallon $16.99")
	fmt.Println(matched, err)

}
