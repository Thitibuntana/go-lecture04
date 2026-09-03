package myarray

import "fmt"

func myArrayA() {
	var data01 [3]int
	data01[0] = 100
	data01[1] = 200
	data01[2] = 300
	fmt.Println(data01[0], data01[1], data01[2])

	var data02 = []int{10, 20, 30}
	fmt.Println(data02[0], data02[1], data02[2])

	data03 := [3]int{1, 2, 3}
	fmt.Println(data03[0], data03[1], data03[2])

	data04 := [...]int{100000, 200000, 300000}

	for i := 0; i < len(data04); i++ {
		fmt.Println(data04[i])
	}

	for i, v := range data04 {
		fmt.Println("i:", i, "v:", v)
	}
}
