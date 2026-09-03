package myarray

import "fmt"

func MyarrayB() {
	data1 := [6]int{1, 2, 3, 4, 5, 6}

	fmt.Println(data1)

	data2 := data1[1:4]
	fmt.Println(data2)

}
