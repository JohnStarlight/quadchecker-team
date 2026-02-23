package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// 1. Read input from the pipe
	inputData, err := io.ReadAll(os.Stdin)
	// ERROR 1: System issue reading the input
	if err != nil {
		fmt.Printf("❌ System Error: Failed to read data (%v)\n", err)
		return
	}

	// ERROR 2: User provided no data
	if len(inputData) == 0 {
		fmt.Println("⚠️  Warning: No data provided.")
		fmt.Println("💡 Tip: Please provide a shape first. Example: ./quadA 3 3 | ./quadchecker")
		return
	}

	input := string(inputData)

	// --- INTERACTIVE CONFIRMATION BLOCK ---
	// Display the parsed shape and ask the user for confirmation.
	// We use /dev/tty to read the user's response (ENTER/q) because
	// os.Stdin is already occupied by the piped shape data.
	fmt.Println("--- RECEIVED SHAPE ---")
	fmt.Print(input)
	fmt.Println()
	fmt.Print("The input forms the shape above. Do you want to proceed? [Press ENTER to match or 'q' to exit]:")

	tty, err := os.Open("/dev/tty")
	if err == nil {
		reader := bufio.NewReader(tty)
		response, _ := reader.ReadString('\n')
		tty.Close()

		if strings.TrimSpace(response) == "q" {
			fmt.Println("Process exited by user.")
			return
		}
	}

	// 2. Check if the input forms a uniform grid and get its dimensions
	x, y, isValid := getDimensions(input)

	// ERROR 3: The input has uneven lines (not a valid grid)
	if !isValid {
		fmt.Println("📐 Format Error: The provided shape has uneven lines (it's not a Quad).")
		fmt.Println("💡 Tip: Make sure all lines of the shape have the exact same length.")
		return
	}

	// 3. List of Quads
	quads := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}
	var matches []string

	safeInput := strings.TrimRight(input, "\n")

	// 4. Check against all known Quads
	for _, name := range quads {
		generated := generateQuad(name, x, y)
		safeGenerated := strings.TrimRight(generated, "\n")

		if safeGenerated == safeInput {
			matches = append(matches, fmt.Sprintf("[%s] [%d] [%d]", name, x, y))
		}
	}

	// ERROR 4: The input is a valid grid, but it doesn't match any known Quad
	if len(matches) == 0 {
		fmt.Printf("🔍 Result: No matching Quad found for this %dx%d input.\n", x, y)
		return
	}

	// SUCCESS! Check if we found one match or multiple matches
	if len(matches) > 1 {
		fmt.Println("✅ Matches found:", strings.Join(matches, " || "))
	} else {
		fmt.Println("✅ Match found:", strings.Join(matches, " || "))
	}
}

// Calculates width (x) and height (y) and checks if the input is a uniform grid
func getDimensions(s string) (int, int, bool) {
	s = strings.TrimSuffix(s, "\n")
	lines := strings.Split(s, "\n")

	y := len(lines)
	if y == 0 {
		return 0, 0, false
	}

	x := len(lines[0])
	if x == 0 {
		return 0, 0, false
	}

	for _, line := range lines {
		if len(line) != x {
			return 0, 0, false
		}
	}

	return x, y, true
}

// Creates the perfect Quad in memory to compare it
func generateQuad(name string, x int, y int) string {
	result := ""
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			char := getChar(name, x, y, i, j)
			result = result + string(char)
		}
		result = result + "\n"
	}
	return result
}

// Returns the correct character depending on the Quad and the position (corners/edges)
func getChar(quad string, x, y, i, j int) byte {
	isTop := (i == 0)
	isBottom := (i == y-1)
	isLeft := (j == 0)
	isRight := (j == x-1)

	if isTop && isLeft {
		switch quad {
		case "quadA":
			return 'o'
		case "quadB":
			return '/'
		case "quadC", "quadD", "quadE":
			return 'A'
		}
	}
	if isTop && isRight {
		switch quad {
		case "quadA":
			return 'o'
		case "quadB":
			return '\\'
		case "quadC":
			return 'A'
		case "quadD", "quadE":
			return 'C'
		}
	}
	if isBottom && isLeft {
		switch quad {
		case "quadA":
			return 'o'
		case "quadB":
			return '\\'
		case "quadC", "quadE":
			return 'C'
		case "quadD":
			return 'A'
		}
	}
	if isBottom && isRight {
		switch quad {
		case "quadA":
			return 'o'
		case "quadB":
			return '/'
		case "quadC", "quadD":
			return 'C'
		case "quadE":
			return 'A'
		}
	}
	if isTop || isBottom {
		switch quad {
		case "quadA":
			return '-'
		case "quadB":
			return '*'
		case "quadC", "quadD", "quadE":
			return 'B'
		}
	}
	if isLeft || isRight {
		switch quad {
		case "quadA":
			return '|'
		case "quadB":
			return '*'
		case "quadC", "quadD", "quadE":
			return 'B'
		}
	}
	return ' '
}
