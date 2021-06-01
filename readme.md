# Golang tutorial

Youtube Link: https://www.youtube.com/watch?v=YS4e4q9oBaU&ab_channel=freeCodeCamp.org

## Project structure

- The main code file is placed in a folder structure "github.com/YellowDemon/" so that whenever you want to fetch this app from an external target, you can use the command "go fetch github.com/YellowDemon/<appName>"

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