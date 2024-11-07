package main

import (
	"fmt"
	"unicode/utf8"
)

/*
NOTE:
1. Bit is a single binary digit, its either 0 or 1
2. Byte is a group of 8 bits, One byte can represent 256
   different value from 0 to 255
3.
*/

func main() {
	var intNum int = 1234
    fmt.Println(intNum)

    // NOTE: Only positive integer

    var floatNum float32 = 432143213421434214321432143214321432
    fmt.Println(floatNum)

    var floatNum32 float32 = 10.1
    var intNum32 int32 = 2
    var result float32 = floatNum32 + float32(intNum32)
    fmt.Println(result)

    var intNum1 int = 3
    var intNum2 int = 2
    fmt.Println(intNum1/intNum2)
    fmt.Println(intNum1%intNum2)

    var myString string = "Hello World"
    var multilineString string = `hello 
world`
    fmt.Println(myString)
    fmt.Println(multilineString)

    fmt.Println(utf8.RuneCountInString("y"))

    var myBoolean bool
    fmt.Println(myBoolean)

    var myvar = "text"
    myvar2 := "text2"
    fmt.Println((myvar))
    fmt.Println(myvar2)

    const myConst string = "const value"
    fmt.Println(myConst)

    printMe()
}

func printMe(){
    fmt.Println("Function")
}