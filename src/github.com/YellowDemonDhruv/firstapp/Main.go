
package main // Has to be named as 'main'

import (
	"fmt" ;
	"strconv"
) 

var A int = 100 // UPPERCASE - Exposed to package global scope

var z byte // Unused variable at package level

var l byte = 50 // Pacakge level variables - we cannot use the ':=' syntax 


// declaring groups of variables together - variable blocks
var ( 
	actorName string = "Dhruv Parthasarathy"
	companion string = "Aditya Ananth"
	doctorNumber int = 3
	season int = 11
)

var ( 
	counter int = 0
)


func main() {

		var i int  // var name type
		i = 42
		fmt.Println(`i = ` , i)

		
		var j int = 50  // var name type = value
		fmt.Println(`j = `, j)

		// DYNAMIC TYPE DECLARATION
		k := 68  // name := value
		fmt.Printf("Type of %v is %T \n", k, k)

		// String formatting
		fmt.Printf("l = %v of type %T \n", l, l )
		
		
		// Type conversion ===================
		b := 42.59
		fmt.Printf("Type of variable b is %T \n", b)
		d := int(b)
		fmt.Printf("Type of variable d is %T \n", d)
		c := strconv.Itoa(d)
		fmt.Printf("Type of variable c is %T \n", c)


	
}


