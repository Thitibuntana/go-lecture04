package myslice

import "fmt"

func mysliceA() {
	var data01 []int
	data01 = append(data01, 10)
	data01 = append(data01, 20)
	fmt.Println(data01[0], data01[1])

	data01 = append(data01, data01[0]+data01[1])
	fmt.Println(data01[0], data01[1], data01[2])

	for i := 0; i < len(data01); i++ {
		fmt.Println(data01[i])
	}

	for i, v := range data01 {
		fmt.Println("i:", i, "v:", v)
	}
}
