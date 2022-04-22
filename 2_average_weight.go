package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Program that calculates the average weight of 5 people.
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter 5 people's weight with spaces: ")
	input_weight, _ := reader.ReadString('\n')

	for i := 0; i < 5; i++ {
		// Split input_weight into array
		input_weight_array := strings.Split(input_weight, " ")
		// Convert string array to int array
		input_weight_array_int := make([]int, len(input_weight_array))

		for j := 0; j < len(input_weight_array); j++ {
			input_weight_array_int[j], _ = strconv.Atoi(input_weight_array[j])
		}

		// Calculate average weight
		average_weight := 0
		for j := 0; j < len(input_weight_array_int); j++ {
			average_weight += input_weight_array_int[j]
		}
		average_weight = average_weight / len(input_weight_array_int)
		fmt.Println("Average weight: ", average_weight)

		// Calculate average
		// average_weight := (input_weight_array[i] + input_weight_array[i+1] + input_weight_array[i+2] + input_weight_array[i+3] + input_weight_array[i+4]) / 5

		// fmt.Println("Average weight: ", average_weight)
	}

	// weight_array_int, _ := strconv.ParseInt(input_weight_array, 10, 64)
	// fmt.Println("You entered: ", weight_array_int)

}
