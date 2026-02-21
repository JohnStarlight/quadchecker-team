package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) > 3 {
		fmt.Println("Error: Too many arguments")
		return
	}
	if len(args) < 3 {
		fmt.Println("Error: Missing arguments")
		return
	}
	x, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Error: 1st argument is not a valid input")
		return
	}
	y, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Error: 2nd argument is not a valid input")
		return
	}

	if x <= 0 {
		fmt.Println("Error: The first argument is invalid. Please input a positive number.")
		return
	}

	if y <= 0 {
		fmt.Println("Error: The second argument is invalid. Please input a positive number.")
		return
	}

	result := ""

	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if row == 0 && col == 0 {
				result += "A"
			} else if row == 0 && col == x-1 {
				result += "C"
			} else if row == y-1 && col == 0 {
				result += "C"
			} else if row == y-1 && col == x-1 {
				result += "A"
			} else if row == 0 || row == y-1 || col == 0 || col == x-1 {
				result += "B"
			} else {
				result += " "
			}
		}
		result += "\n"
	}
	fmt.Print(result)
}
