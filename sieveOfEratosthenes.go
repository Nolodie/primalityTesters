package main

import "fmt"

func main() {

	//Create an array of numbers.

	var numbs[49] int
	u := 2
	for j := 0; j < len(numbs); j++ {
		numbs[j] = u
		u++
	}

	//Sieve for primes.

	for i := 0; i < len(numbs); i++ {
		var factor int
		factor = numbs[i]
	//*	fmt.Printf("%d\n", factor)	//check 1: Print factor/prime.
		if factor != 0 {
			for k := 1; k < len(numbs); k++ {
				var check int
				check = numbs[k]
				if check % factor == 0 {
	//*				fmt.Printf("%d ", check)	//check 2: Print factorized numbers.
					numbs[k] = 0
				}
			}
	//*		fmt.Printf("\n")	//check 3: Formating for previous two checks.

			//Print primes.

			fmt.Printf("%d ", factor)
		}
	}
	fmt.Printf("\n")
}


