package main

import (
	"calculation/calculation"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Checking if the correct number of arguments is provided
	if len(os.Args) == 1 {
		fmt.Println("File name missing")
		return
	} else if len(os.Args) != 2 {
		fmt.Println("Wrong use of the tool")
		return
	}
	// Get the input file name from command line arguments
	inputfile := os.Args[1]
	// Read the contents of the file
	input, err := os.ReadFile(inputfile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Split file contents into lines
	lines := strings.Split(string(input), "\n")
	lines = lines[:len(lines)-1]
	var data []float64
	for _, line := range lines {
		if line == "" {
			continue
		}
		value, err := strconv.ParseFloat(line, 64) // Convert string to float64
		if err != nil {
			fmt.Printf("Error converting line '%s' to float: %v\n", line, err)
			return
		}
		data = append(data, value) // Append the converted value to data slice
	}

	Statistic(data)
}

// Print various statistics calculated from data
func Statistic(data []float64) {
	fmt.Println("Average:", int(math.Round(calculation.Average(data))))
	fmt.Println("Median:", int(math.Round(calculation.Median(data))))
	fmt.Println("Variance:", int(math.Round(calculation.Variance(data))))
	fmt.Println("Standard Deviation:", int(math.Round(calculation.Deviation(data))))
}
