package mypointer

import "fmt"

func MyPointerA() {
	data := 20
	fmt.Println("data:", data)

	var p *int
	p = &data
	fmt.Println("p:", p)

	*p = 30
	fmt.Println("data:", data)

	data02 := 40
	fmt.Println("data02:", data02)

	p = &data02
	fmt.Println("p:", p)

	*p = 50
	fmt.Println("data02:", data02)
}
