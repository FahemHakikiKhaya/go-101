package main

import "fmt"

func main() {
	var intArr [3]int32
	intArr[0] = 1
	intArr[1] = 2
	fmt.Println(intArr[0])
	fmt.Println(intArr[0:2])

	var intArr2 [3]int32 = [3]int32{1,2,3}
	fmt.Println(intArr2)

	var intSlice []int32 = []int32{4,5,6}
	fmt.Printf("The length is %v with capacity %v",len(intSlice),cap(intSlice))

	intSlice = append(intSlice, 7)
	fmt.Printf("\nThe length is %v with capacity %v",len(intSlice),cap(intSlice))


	var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice,intSlice2...)

	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)
}