package main

import "fmt"

func testCount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testCount(x + 1)
}

func main() {
	//var arr1 = [...]int{1, 2, 3, 4, 5, 6}
	//slice := arr1[2:3]
	//fmt.Printf("slice = %v\n", slice)
	//fmt.Printf("length = %d\n", len(slice))
	//fmt.Printf("capacity = %v\n", cap(slice))
	testCount(1)
}
