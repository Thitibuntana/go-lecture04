package mystruck

type SauStudent struct {
	id     string
	name   string
	gender string
}

func MySauStudent() {
	var s [3]SauStudent
	s[0] = SauStudent{
		id:     "670601",
		name:   "Tonkla",
		gender: "Male",
	}
	s[1] = SauStudent{
		id:     "670602",
		name:   "Phatchara",
		gender: "Female",
	}
	s[2] = SauStudent{
		id:     "670603",
		name:   "Nawarat",
		gender: "Female",
	}

}
