package quads

func QuadA(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}

	result := ""

	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if (row == 0 || row == y-1) && (col == 0 || col == x-1) {
				result += "o"
			} else if row == 0 || row == y-1 {
				result += "-"
			} else if col == 0 || col == x-1 {
				result += "|"
			} else {
				result += " "
			}
		}
		result += "\n"
	}

	return result
}
