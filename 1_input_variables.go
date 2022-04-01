package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	// Create Time & Format output
	current_date := time.Now().Format("2006-1-02")
	// Split Time into date strings array
	date_strings := strings.Split(current_date, "-")
	// Split and convert date into int
	current_year, _ := strconv.ParseInt(date_strings[0], 10, 64)
	current_month, _ := strconv.ParseInt(date_strings[1], 10, 64)
	current_day, _ := strconv.ParseInt(date_strings[2], 10, 64)

	fmt.Println(current_year, current_month, current_day)

	fmt.Print("\n")
	fmt.Println("Hello, \nWhat's your age?")
	user_age, _ := reader.ReadString('\n')
	// ??? Convert String to Int
	int_user_age, _ := strconv.ParseInt(user_age, 10, 64)

	fmt.Printf("%d", int_user_age)
	fmt.Print("\n")

	fmt.Println("What month were you born in?")
	user_month, _ := reader.ReadString('\n')
	// If input === month, mon, convert to #
	// fmt.Printf("%02d, %02d, %02d, %02d", jan, dec, july, nov)
	fmt.Printf("What day of %s? ", strings.TrimSpace(user_month))
	user_day, _ := reader.ReadString('\n')

	// birth_year := (current_year - age)

	// Remove this line
	// fmt.Printf("Your birthdate is %s %s %s", birth_year, strings.TrimSpace(user_month), user_day)
	fmt.Printf("Your birthdate is %s %s %s", strings.TrimSpace(user_age), strings.TrimSpace(user_month), user_day)
	// Take this string & measure against 'current_date'

	birth_date := strings.TrimSpace(user_age) + "-" + strings.TrimSpace(user_month) + "-" + user_day
	fmt.Println(birth_date)

	// Return "DOB YYYY-mm-dd"

}
