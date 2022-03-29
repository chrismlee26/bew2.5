package main

import (
	"bufio"
	"fmt"
	"os"

	// "time"
	"strings"
)

// 1. Write a program that calculates the birth year using a provided date of birth and age. HINT: Get the date of birth and age from stdin!
// 2. Write a program that calculates the average weight of 5 people.

// var (
// 	day   int
// 	month int
// 	year  int
// )

// const (
// 	jan = iota + 1
// 	feb
// 	mar
// 	apr
// 	may
// 	jun
// 	jul
// 	aug
// 	sep
// 	oct
// 	nov
// 	dec
// )

// const (
// 	january = iota + 1
// 	february
// 	march
// 	april
// 	june
// 	july
// 	august
// 	september
// 	october
// 	november
// 	december
// )

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Hello, \nWhat year were you born in?")
	year, _ := reader.ReadString('\n')
	fmt.Println("What month were you born in?")
	month, _ := reader.ReadString('\n')
	fmt.Printf("What day of %s? ", strings.TrimSpace(month))
	day, _ := reader.ReadString('\n')
	fmt.Printf("Your birthdate is %s %s %s", strings.TrimSpace(day), strings.TrimSpace(month), year)
}
