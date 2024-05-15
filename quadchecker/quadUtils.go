package quadchecker

func QuadA(x, y int) []string {
	if x <= 0 || y <= 0 {
		return nil
	}

	// Create a slice to hold the lines
	lines := make([]string, 0, y)

	// Loop over each row
	for i := 0; i < y; i++ {
		line := ""
		// Loop over each column
		for j := 0; j < x; j++ {
			// Append the appropriate character based on position
			if i == 0 && j == 0 || i == 0 && j == x-1 || i == y-1 && j == 0 || i == y-1 && j == x-1 {
				line += "o"
			} else if i == 0 || i == y-1 {
				line += "-"
			} else if j == 0 || j == x-1 {
				line += "|"
			} else {
				line += " "
			}
		}
		// Append the line to the slice
		lines = append(lines, line)
	}

	return lines
}

func QuadB(x, y int) []string {
	if x <= 0 || y <= 0 {
		return nil
	}

	// Create a slice to hold the lines
	lines := make([]string, 0, y)

	// Loop over each row
	for i := 0; i < y; i++ {
		line := ""
		// Loop over each column
		for j := 0; j < x; j++ {
			// Append the appropriate character based on position
			if i == 0 && j == 0 {
				line += "/"
			} else if i == 0 && j == x-1 {
				line += "\\"
			} else if i == y-1 && j == 0 {
				line += "\\"
			} else if i == y-1 && j == x-1 {
				line += "/"
			} else if i == 0 || i == y-1 {
				line += "*"
			} else if j == 0 || j == x-1 {
				line += "*"
			} else {
				line += " "
			}
		}
		// Append the line to the slice
		lines = append(lines, line)
	}

	return lines
}

func QuadC(x, y int) []string {
	if x <= 0 || y <= 0 {
		return nil
	}

	// Create a slice to hold the lines
	lines := make([]string, 0, y)

	// Loop over each row
	for i := 0; i < y; i++ {
		line := ""
		// Loop over each column
		for j := 0; j < x; j++ {
			// Append the appropriate character based on position
			if i == 0 && j == 0 {
				line += "A"
			} else if i == 0 && j == x-1 {
				line += "A"
			} else if i == y-1 && j == 0 {
				line += "C"
			} else if i == y-1 && j == x-1 {
				line += "C"
			} else if i == 0 || i == y-1 {
				line += "B"
			} else if j == 0 || j == x-1 {
				line += "B"
			} else {
				line += " "
			}
		}
		// Append the line to the slice
		lines = append(lines, line)
	}

	return lines
}

func QuadD(x, y int) []string {
	if x <= 0 || y <= 0 {
		return nil
	}

	// Create a slice to hold the lines
	lines := make([]string, 0, y)

	// Loop over each row
	for i := 0; i < y; i++ {
		line := ""
		// Loop over each column
		for j := 0; j < x; j++ {
			// Append the appropriate character based on position
			if i == 0 && j == 0 {
				line += "A"
			} else if i == 0 && j == x-1 {
				line += "C"
			} else if i == y-1 && j == 0 {
				line += "A"
			} else if i == y-1 && j == x-1 {
				line += "C"
			} else if i == 0 || i == y-1 {
				line += "B"
			} else if j == 0 || j == x-1 {
				line += "B"
			} else {
				line += " "
			}
		}
		// Append the line to the slice
		lines = append(lines, line)
	}

	return lines
}

func QuadE(x, y int) []string {
	if x <= 0 || y <= 0 {
		return nil
	}

	// Create a slice to hold the lines
	lines := make([]string, 0, y)

	// Loop over each row
	for i := 0; i < y; i++ {
		line := ""
		// Loop over each column
		for j := 0; j < x; j++ {
			// Append the appropriate character based on position
			if i == 0 && j == 0 {
				line += "A"
			} else if i == 0 && j == x-1 {
				line += "C"
			} else if i == y-1 && j == 0 {
				line += "C"
			} else if i == y-1 && j == x-1 {
				line += "A"
			} else if i == 0 || i == y-1 {
				line += "B"
			} else if j == 0 || j == x-1 {
				line += "B"
			} else {
				line += " "
			}
		}
		// Append the line to the slice
		lines = append(lines, line)
	}

	return lines
}
