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

// SWITCH STATEMENTS

fmt.Println("\nSWITCH CASE")

	switch 3 {
	case 1, 5, 10:
		fmt.Println("one, five, or 10")
	case 2, 4, 6:
		fmt.Println("two, four or 6")
	default:
		fmt.Println("another number")
	}

	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one, five, or 10")
	case 2, 4, 6:
		fmt.Println("two, four or 6")
	default:
		fmt.Println("another number")
	}

	i:= 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to 10")
	case i <= 20:
		fmt.Println("less than or rqual to 20")
	default:
		fmt.Println("greater than 20")
	}


// Fall through - continues to execute the next statement without breaking

	switch {
	case i <= 10:
		fmt.Println("less than or equal to 10")
		fallthrough
	case i <= 20:
		fmt.Println("less than or rqual to 20")
	default:
		fmt.Println("greater than 20")
	}

	var j interface{} = [4]string{}
	switch j.(type) {
	case int :
		fmt.Println("j is an int")
	case float64:
		fmt.Println("j is a float64")
	case [4]int:
		fmt.Println("j is a [4]int object")
		if true{
			break
		}
		fmt.Println("This line is not executed as code breaks out of switch block")
	default:
		fmt.Println("Idk the type of this object")
	}

}
