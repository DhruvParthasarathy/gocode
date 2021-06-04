package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	if (true) {
		fmt.Println("Inside true block")
	}

	statePopulations := map[string]int {
		"California" : 123,
	}

	if pop, ok := statePopulations["California"]; ok {
		fmt.Println("California pop: ",pop)
	}


// Number guessing game
	
	number:= 50
	guess := 50 // Change the guess to see which condition gets triggered

	if guess < number {
		fmt.Println("Too low")
	} else if  guess > number {
		fmt.Println("Too high")
	} else {
		fmt.Println("Got it!")
	}

// Other arithmetic comparison operators are <=, >= and !=

// Logical operators '||', '&&', and '!'

	// OR Operator
	if guess < 1 || guess > 100 {
		fmt.Println("The guess must be between 1 and 100!")
	}

	// AND Operator
	if ( (guess >= 1) && (guess <= 100) ){

		fmt.Println("Guess is within the range")
	}

	// ! Operator
	fmt.Println(!true)

	if guess < 1 || guess > 100 {
		fmt.Println("Out of bounds")
	} else {
		fmt.Println("Within bounds")
	}



}
