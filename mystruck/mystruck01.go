package mystruck

type Mobile struct {
	name  string
	brand string
	price float64
}

func MyMobile() {
	m1 := Mobile{
		name:  "iPhone 15",
		brand: "Apple",
		price: 45000,
	}

	println(m1.name)
	println(m1.brand)
	println(m1.price)

	m2 := Mobile{}
	m2.name = "Galaxy S24"
	m2.brand = "Samsung"
	m2.price = 35000

	println(m2.name)
	println(m2.brand)
	println(m2.price)

}
