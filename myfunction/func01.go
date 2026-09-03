package myfunction

import "fmt"

//no parameter, no return value
func doA() {
	fmt.Println("Hello")
}

//have parameter, no return value
func doB(x int, y int) {
	fmt.Println(x + y)
}

//no parameter, have return value
func doC() int {
	fmt.Println("Hello")
	return 100
}

//have parameter, have return value
func doD() (int, string) {
	fmt.Println("Hello")
	return 555, "Hello"
}

func doE(p1 int, p2 string, p3 bool) (int, string, bool) {
	fmt.Println("Hello")
	return 100, "Hello", true

}
