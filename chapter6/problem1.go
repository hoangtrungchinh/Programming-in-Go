// chapter 5: control structers
package main

import "fmt"

func main() {
	// problem 2
	// Write a program that prints out all the numbers
	// evenly divisible by 3 between 1 and 100. (3, 6, 9,
	// etc.)
	for i := 2; i < 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}

	// problem 3
	// Write a program that prints the numbers from 1
	// to 100. But for multiples of three print "Fizz" instead of the number and for the multiples of five
	// print "Buzz". For numbers which are multiples
	// of both three and five print "FizzBuzz"
	for i := 2; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
