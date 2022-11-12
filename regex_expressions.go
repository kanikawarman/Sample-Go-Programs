//Various examples of regex expressions - Go grammar rules
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
	/*Write regex expressions for the Dimensions of most products on the Home Depot website, with H for height, W for width, and D for depth.
	Ex: H 32.7 in, W 80.3 in, D 31.5 in */
	matched, err := regexp.MatchString(`^(H .*\..* in, W .*\..* in, D .*\..* in)$`, "H 32.7 in, W 80.3 in, D 31.5 in")
	fmt.Println(matched, err)
	/* And for A citation for lines from a play. Ex: Six Degrees of Separation 1.3.188-90.
	If there is more than one line, it is specified by the first line, a dash (-) and then the last line.*/
	matched, err = regexp.MatchString(`^(.*\..*\..*)$`, "Six Degrees of Separation 1.3.188-90")
	fmt.Println(matched, err)

}
