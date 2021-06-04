# Golang tutorial

*Hi there! this document contains key snippets of information that I have been able to make note of while going through a 6 hour video on the basics of Golang, the from the following [Youtube Video ](https://www.youtube.com/watch?v=YS4e4q9oBaU&ab_channel=freeCodeCamp.org).  '*

*Before starting off with exploring the code, you might want to have a look at the following [link](https://www.cyberciti.biz/faq/how-to-install-gol-ang-on-ubuntu-linux/) or follow the steps mentioned in the first part of the above video to set up Golang in your machine'*

*Once you are done setting up Golang, to quickly start off you can 'cd' to 'src/github.com/YellowDemonDhruv/firstapp/' and run the command go run Main.go'*

## Project structure

- The main code file is placed in a folder structure "github.com/YellowDemonDhruv/" so that whenever you want to fetch this app from an external target, you can use the command "go fetch github.com/YellowDemonDhruv/gocode"

- When we build a go file, using go build <filePath>, the output is a compiled binary file which can be directly run on the terminal by just typing in the file name

## Variables


### Variable definition in GOLANG

- An UPPERCASE variable is exposed to the global package scope
- Variables declared at the package level 'can be not used' throughout the program, but variables declared in any local scope of a function have to be used
- While declating pacakge level variables - we cannot use the ':=' syntax here and we have to be explicit with the keyword var and the variable type

### Type conversion in GOLANG

- while converting an integer to a type string - go just checks what is the character which contains the unicode value of that integer, For example, if you declare i = 42 and set j = string(i), then you get * as the value of j, because the character * has the unicode value of 42. So to achieve conversion between numbers and strings, we have a package in golang called as 'strconv'.
- For conversion between numeric types we need to use explicit type conversion. 

- ![Variables basics](img/variables.png)
- ![Naming convenstions and type conversions](img/naming_types.png)


### Primitives

#### Complex numbers are considered as a first class citizen in golang - there is a complex32 and complex64 data type in golang

#### Boolean type
- Values are true or false
- Not an alias for other types(e.g. int)
- Zero value is false

#### Numeric types

Integers
- Signed integers ( both positive and negative integers )
    - int type has varying size, but min 32 bits ( the most common type of integer that we would be dealing in any application )
    - If we need more control we can use the 8bit(int8) through 64 bit(int64)

- Unsigned integers ( only positive integers )
    - 8 bit (byte and uint8 ) through 32 bit (unit32)

- Arithmetic operatons
    - Addition + , subtraction - , multiplication * , division / , remainder %

- Bitwise operations
    - And '&', or '|', xor '^' ,and not '&^'

- Can't mix types in the same family! (uint16 + uint32 = error)


Floating point
- 32 and 64 bit versions

- Literal styles
    - Decimal(3.14)
    - Exponential(12e18 or 12E10)
    - Mixed(12.5e15)

- Arithmetic operations
    - Add, sub, mul, div, <s>remainder</s>

Complex number
- Zero value is 0 + 0i

- Built up from float numbers ( 32 or 64 ) 

- These numbers have a 64 or 128 bit versions based on what they are built up of

- Built in functions
    - complex() - make complex number(64, 128) from 2 floats(32, 64)
    - real() - get the real part as a float
    - imag() - get the imaginary part as a float

#### Text types

Strings
- UTF-8
- Immutable
- Can be concatenated with the + operator
- Can be converted to []byte

Rune
- UTF-32
- ALias for int32
- Special methods normally required to process
    - E.g. [strings.Reader#ReadRune](https://golang.org/pkg/strings/#Reader.ReadRune)


## Constants

### Constants can only be set as values which can be read during compile time

Eg: If you want to set the value of a constant c as math.Sin(1.57), this cannot happen because, to set the value of c, the fucntion math.Sin() has to be executed. 

Constants can be made of any other primitive types as the variables

### Constants are super cool !!
Constants need not be strictly typed. For example, if you declate var i = 10, the type of i will always be an 'int' unless it is strictly typed to be of any other type. But this is not the case with constants. If we declare const a = 10, we can perform an addition with any int, float or even a complex variable or constant.

- Constants are immutable, but can be shadowed - we can not only change the value but also the type 

- The value of the constant has to be calculable at compile time

 - Named like variables
    - PascalCase for exported constants
    - camelCase for internal constants

- Typed constants work like immutable variables
    - Can interoperate only with the same type 

- Untyped constants work like literals
    - Can interoperate with similar types : Constants are super cool !!

## Enumerated constants

- Special symbol iota allows related constants to be created easily
- Iota starts at 0 in each const block and increments by one
- Watch out for constant values that match zero values for variables

## Enumerated expressions
- Operations that can be determined at compile time are allowed
    - Arithmetic
    - Bitwise operations
    - Bitshifting

## Arrays 

- Collecton of items with the same type

- Fixed size at compile time

- Declaration styles
    - a := [3]int{1,2,3}
    - a := [...]int{1,2,3}
    - var a [3]int
    
- Access via a zero-based index
    - a := [3]int {1,2,3} // a[1] == 3

- len function returns the size of the array

- Copies refer to different underlynig data, no memory reference, expensive copy operations

## Slices

- Copies of slices refer to same underlying array

### Creation styles
- Slice existing array or slice

- Literal style

- Via make function
    - a := make( []int, 10 ) // create slice with capacity and length == 10

    - a := make( []int, 10, 100 ) // slice with length == 10 and calacity == 100

    - len function returns length of slice ( number of elements present )

    - cap function returns the length of the underlying array

    - append function to add elements to a slice
        - May cause expensive copy operation if underlying array is too small as compared to the number of elements being added, since the slice capacity needs to be reset

## Maps
- Maps are collections of value types that are accessed via keys

```
	mapA := map[string]int {
		"California" : 123,
		"Berkley" : 345,
		"New jersey" : 456,
	}

```
- Created via literals or via the "make" function

- Members are accessed via [key] syntax
    - myMap["key"] = "value"

- Check for presence with "value, ok" form of result
    - pop, ok := statePopulations["Oho"]
    - ok == false

- Multiple assignments refer to the same underlying data, maps are memory referenced

## Structs
- Collections of disparate data types that describe a single concept

- Keyed by named fields

- Normally created as types, but anonymous types are allowed.
    - Anonymous structs are usually short lived

- Structs are value referenced

- No inheritence, but we can use composition via embedding
    ```

    type human struct {
        name string
        age int
    }

    type Doctor struct {
        human // embedded struct
        degree string
        speciality string
    }
    ```

- Tags can be added to struct fields to describe a field

## If statements

- Initializer syntax

    ``` 
    if pop, ok := statePopulations["California"]; ok {
		fmt.Println("California pop: ",pop)
	}
    ```

    The variable 'pop' is scoped only to this if block

- Comparison operators

- Logical operators

- Short circuiting - in an OR operations, if any one of the checks are true, the remaining tests are not going to be checked

- If - else if - else statements

- Equality of floats
    While checking 2 floating point numbers, it is better to check if the error between the 2 numbers is below a certain threshold value rather than directly checking the numbers themselves ( maybe < 0.01 )

## Switch statements

- Switch on a tag

- Cases with multiple tests
    ```
    switch 3 {
	case 1, 5, 10:
		fmt.Println("one, five, or 10")
	case 2, 4, 6:
		fmt.Println("two, four or 6")
	default:
		fmt.Println("another number")
	}
    ```

- Initializers
    ```
    switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one, five, or 10")
	case 2, 4, 6:
		fmt.Println("two, four or 6")
	default:
		fmt.Println("another number")
	}
    ```

- Switches with no tags
    ```
    i:= 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to 10")
	case i <= 20:
		fmt.Println("less than or rqual to 20")
	default:
		fmt.Println("greater than 20")
	}
    ```

- Fallthrough vs break keywords
    - implicit breaks
    - explicit fallthrough

- Type switches 
    - Used on an empty interface

- Breaking out of execution using break statement
    ```
    var j interface{} = [4]string{}
	switch j.(type) {
	case int :
		fmt.Println("j is an int")
	case float64:
		fmt.Println("j is a float64")
	case [4]int:
		fmt.Println("j is a [4]int object")
		if true{
			break // If a condition is satisfied, break out of the statement and do not execute the remaining statements under this case
		}
		fmt.Println("This line is not executed as code breaks out of switch block")
	default:
		fmt.Println("Idk the type of this object")
	}
    ```