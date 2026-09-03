package myslice

import "fmt"

func mysliceC() {

	data1 := []string{"a", "b", "c", "d", "e"}
	fmt.Println(data1)

	data2 := data1[1:4]
	fmt.Println(data2)

	data3 := append(data2, "f")
	fmt.Println(data3)

}
