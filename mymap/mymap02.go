package mymap

func MyMap1() {

	data1 := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"ten":   "10",
	}

	println()
	for _, val := range data1 {
		println(val, " ")
	}
	println()
	for key, val := range data1 {
		println(key, " : ", val)
	}

}
