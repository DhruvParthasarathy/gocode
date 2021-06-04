
package main // Has to be named as 'main'

import (
	// "fmt" ;
	// "strconv"
) 

// var A int = 100 // UPPERCASE - Exposed to package global scope

// var z byte // Unused variable at package level

// var l byte = 50 // Pacakge level variables - we cannot use the ':=' syntax 


// declaring groups of variables together - variable blocks
// var ( 
// 	actorName string = "Dhruv Parthasarathy"
// 	companion string = "Aditya Ananth"
// 	doctorNumber int = 3
// 	season int = 11
// )

// var ( 
// 	counter int = 0
// )


// func main() {

// 		fmt.Println("Working with number primitives ================")
// 		var i int  // var name type
// 		i = 42
// 		fmt.Println(`i = ` , i)

		
// 		var j int = 50  // var name type = value
// 		fmt.Println(`j = `, j)

// 	// DYNAMIC TYPE DECLARATION
// 		k := 68  // name := value
// 		fmt.Printf("Type of %v is %T \n", k, k)

// 	// String formatting
// 		fmt.Printf("l = %v of type %T \n", l, l )
		
		
// 	// Type conversion ===================
// 		b := 42.59
// 		fmt.Printf("Type of variable b is %T \n", b)
// 		d := int(b)
// 		fmt.Printf("Type of variable d is %T \n", d)
// 		c := strconv.Itoa(d)
// 		fmt.Printf("Type of variable c is %T \n", c)

// 		fmt.Println("\nBooleans ================")
// 		n := false
// 		fmt.Printf("%v is of type %T\n", n, n)

// 		// testing equivalence
// 		m := 1 == 1
// 		fmt.Printf("%v is of type %T \n", m, m)

// 	// The zero value of a variable
// 		var (
// 			r int
// 			t bool
// 			h string
// 			q float32
// 		)

		
// 		fmt.Println("\nZero values  ================ ")
// 		fmt.Printf("The values of r, t, h and q are %v, %v, %v, %v \n", r,t,h,q)

// 		fmt.Println("\nThe following bitwise operators are available in golang: &, |, ^ and &^. These are AND, OR, XOR and AND NOR operators respectively")

// 	// Primitives

// 		// COMPLEX NUMBERS
// 			fmt.Println("\nWorkig with complex numbers")
// 			u := 1 + 2i
// 			f := 2 + 5i
// 			fmt.Println(u + f)
// 			fmt.Println(u - f)
// 			fmt.Println(u * f)
// 			fmt.Println(u / f)
// 			fmt.Printf("real() and the imag() function. \nThe value of u is %v, %T. The real coponent of u is %v, %T and the imaginary component of u is %v, %T", u ,u , real(u) ,real(u), imag(u), imag(u))

// 			// The complex() function, this takes in 2 numbers and creates a complex number as the output

// 			var a1 float32 = float32(5)
// 			var a2 float32 = float32(12)
// 			a3  := complex(a1, a2)
// 			fmt.Println("\n \nCreating a complex number from 2 float values")
// 			fmt.Printf("Value of a1 is %v, %T. The value of a2 is %v, %T and the value of a3 is %v, %T", a1, a1, a2, a2, a3, a3)
		
// 		// STRINGS
// 			fmt.Println("\n \nWorking with strings")
// 			s := "this is a string"
// 			fmt.Printf(" \n The value of the string s is ' %v ' ", s)
// 			fmt.Printf(" \n The 6th letter in the string is '%v', %T", string(s[6]), s[6])

// 			fmt.Println("\n \nThis string can be converted to a collection of bytes")
// 			byt := []byte(s)
// 			fmt.Printf("%v, %T\n", byt, byt)


// 		// A RUNE
// 			fmt.Println("\nA rune is a int32 character")
// 			rn := 'a'
// 			fmt.Printf("The value of the rune rn is %v, %T", rn, rn)
		

	
// }


 