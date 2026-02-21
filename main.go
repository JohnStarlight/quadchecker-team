package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// 1. Διαβάζουμε το input από το pipe
	inputData, err := io.ReadAll(os.Stdin)
	if err != nil || len(inputData) == 0 {
		fmt.Println("Not a quad function")
		return
	}
	input := string(inputData)

	// 2. Ελέγχουμε αν είναι σωστό παραλληλόγραμμο και παίρνουμε διαστάσεις
	x, y, isValid := getDimensions(input)
	if !isValid {
		fmt.Println("Not a quad function")
		return
	}

	// 3. Λίστα με τα Quads
	quads := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}
	var matches []string

	// Αφαιρούμε το τελευταίο \n από το input (Παγίδα του Test Script)
	safeInput := strings.TrimRight(input, "\n")

	// 4. Ελέγχουμε ένα-ένα τα Quads
	for _, name := range quads {
		generated := generateQuad(name, x, y)
		safeGenerated := strings.TrimRight(generated, "\n")

		if safeGenerated == safeInput {
			matches = append(matches, fmt.Sprintf("[%s] [%d] [%d]", name, x, y))
		}
	}

	// 5. Αν δεν βρέθηκε κανένα
	if len(matches) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	// 6. Τυπώνουμε τα matches ενωμένα με ||
	fmt.Println(strings.Join(matches, " || "))
}

// Υπολογίζει το πλάτος (x) και ύψος (y) και ελέγχει αν το σχήμα είναι ομοιόμορφο
func getDimensions(s string) (int, int, bool) {
	// 1. Πριν κάνουμε ΟΤΙΔΗΠΟΤΕ, κόβουμε το '\n' που μπαίνει αυτόματα στο τέλος
	s = strings.TrimSuffix(s, "\n")

	// 2. Χωρίζουμε το σχήμα στις γραμμές του
	lines := strings.Split(s, "\n")

	// 3. Βρίσκουμε το ύψος (πόσες γραμμές βγήκαν;)
	y := len(lines)
	if y == 0 {
		return 0, 0, false
	}

	// 4. Βρίσκουμε το πλάτος (πόσους χαρακτήρες έχει η πρώτη γραμμή;)
	x := len(lines[0])
	if x == 0 {
		return 0, 0, false
	}

	// 5. Ελέγχουμε αν όλες οι γραμμές είναι ίσες (τέλειο ορθογώνιο)
	for _, line := range lines {
		if len(line) != x {
			return 0, 0, false // Αν βρει στραβή γραμμή, σταματάει!
		}
	}

	// Αν έφτασε ως εδώ, το σχήμα είναι τέλειο!
	return x, y, true
}

// Δημιουργεί στη μνήμη το τέλειο Quad για να το συγκρίνουμε
func generateQuad(name string, x, y int) string {
	var sb strings.Builder
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			sb.WriteByte(getChar(name, x, y, i, j))
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

// Επιστρέφει τον σωστό χαρακτήρα ανάλογα με το Quad και τη θέση (γωνίες/πλευρές)
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
