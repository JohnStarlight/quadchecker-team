package quads

func QuadC(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}

	result := ""

	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if row == 0 && col == 0 {
				result += "A"
			} else if row == 0 && col == x-1 {
				result += "A"
			} else if row == y-1 && col == 0 {
				result += "C"
			} else if row == y-1 && col == x-1 {
				result += "C"
			} else if row == 0 || row == y-1 || col == 0 || col == x-1 {
				result += "B"
			} else {
				result += " "
			}
		}
		result += "\n"
	}

	return result
}
