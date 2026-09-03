package mystruck

import "fmt"

type DtiStudent struct {
	id   string
	name string
	age  int
}

func MyDtiStudent() {
	s1 := DtiStudent{
		id:   "670601",
		name: "Tonkla",
		age:  20,
	}

	fmt.Println(s1.id, s1.name, s1.age)

	s2 := DtiStudent{}
	s2.id = "670602"
	s2.name = "Phatchara"
	s2.age = 19

	fmt.Println(s2.id, s2.name, s2.age)

	//struck in short term
	s3 := DtiStudent{"670603", "Nawarat", 19}
	fmt.Println(s3.id, s3.name, s3.age)

	s4 := DtiStudent{
		id:   "670604",
		name: "Nawarat",
		age:  19,
	}
	fmt.Println(s4.id, s4.name, s4.age)

}
