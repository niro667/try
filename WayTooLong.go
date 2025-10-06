package main

import (
	"fmt"
	"strconv"
)

func main() {

	/*i want to get a user input to run slice
	every word he enters i count the charachters of it so if the word > 10 i will print the first and the last latter and all the numbers of the charachters
	between them */

	var choice int8
	var first string

	wordz := []string{}
	fmt.Scan(&choice)
	for b := 0; b < int(choice); b++ {

		fmt.Scan(&first)
		wordz = append(wordz, first)
	}
	var add = 0

	firstl, lastl := "", ""
	for _, val := range wordz {
		if len(val) > 10 {
			for i := 0; i < len(val)-1; i++ {
				add = i
				firstl, lastl = string(val[0]), string(val[len(val)-1])
			}
			out := firstl + strconv.Itoa(add) + lastl
			fmt.Println(out)

		} else {
			fmt.Println(val)
		}

	}

}
