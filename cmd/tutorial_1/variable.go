package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// ==============================
	// SECTION 1: Variables and Basic Data Types
	// ==============================

	// Integer
	var intNum int = 1234
	fmt.Println(intNum)

	// Float
	var floatNum float32 = 432143213421434214321432143214321432
	fmt.Println(floatNum)

	// Integer Division
	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	// Strings
	var myString string = "Hello World"
	var multilineString string = `hello world`
	fmt.Println(myString)
	fmt.Println(multilineString)

	// Unicode
	fmt.Println(utf8.RuneCountInString("y"))

	// Boolean
	var myBoolean bool
	fmt.Println(myBoolean)

	// Variable Declaration
	var myvar = "text"
	myvar2 := "text2"
	fmt.Println(myvar)
	fmt.Println(myvar2)

	// Constants
	const myConst string = "const value"
	fmt.Println(myConst)
}