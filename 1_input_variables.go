package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// 1. Write a program that calculates the birth year using a provided date of birth and age. HINT: Get the date of birth and age from stdin!
// 2. Write a program that calculates the average weight of 5 people.

// var (
// 	day   int
// 	month int
// 	year  int
// )

const (
	jan = iota + 1
	feb
	mar
	apr
	may
	jun
	jul
	aug
	sep
	oct
	nov
	dec
)

const (
	january   = jan
	february  = feb
	march     = mar
	april     = apr
	june      = jun
	july      = jul
	august    = aug
	september = sep
	october   = oct
	november  = nov
	december  = dec
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	// Time Library
	current_date := time.Now().Format("2006-1-02")
	// fmt.Println(current_date.Format("2006-1-02"), '\n')
	fmt.Println(current_date, "~~~~~~~~~~~")

	// Testing Variables
	fmt.Printf("%02d, %02d, %02d, %02d", jan, dec, july, nov)
	// ---------------
	// Convert time to string & split
	// current_date_string := current_date.String()
	// current_year := strings.Split(current_date_string, "-")
	// fmt.Println(current_year)

	fmt.Println("\n Hello, \nWhat year were you born in?")
	year, _ := reader.ReadString('\n')
	fmt.Println("What month were you born in?")
	month, _ := reader.ReadString('\n')
	fmt.Printf("What day of %s? ", strings.TrimSpace(month))
	day, _ := reader.ReadString('\n')
	// Remove this line
	fmt.Printf("Your birthdate is %s %s %s", strings.TrimSpace(year), strings.TrimSpace(month), day)
	// Take this string & measure against 'current_date'
	birth_date := strings.TrimSpace(year) + "-" + strings.TrimSpace(month) + "-" + day
	fmt.Println(birth_date)

	// Return "age"
}
