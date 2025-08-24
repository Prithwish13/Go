package basics


func switchIfCase() {
	a := 29

	if a >= 18 && a <= 21 {
		println("You are an adult")
	} else if a > 21 && a < 25 {
		println("You are on a pick state of your life")		
	} else {
		println("This is time for struggle! Keep going!")
	}

	// Switch case example
	switch a {
	case 18:
		println("You are 18")
	case 19:
		println("You are 19")
	case 20:
		println("You are 20")
	case 21:
		println("You are 21")
	case 22:
		println("You are 22")
	case 23, 24, 25, 26, 27, 28, 29:
		println("You are in a range of 23 to 29")
	default:
		println("You are not in the range of 18 to 22")
	}

	// Switch case with fallthrough
	switch a {
	case 18:
		println("You are 18")
		fallthrough // This will execute the next case as well
	case 19:
		println("You are 19")
	case 20:
		println("You are 20")
	default:
		println("You are not in the range of 18 to 20")
	}

	// Type switch example

	checkType(42)
}

func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		println("Integer:", v)
	case string:
		println("String:", v)
	case bool:
		println("Boolean:", v)
	default:
		println("Unknown type")

	}
}