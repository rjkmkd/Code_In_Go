package main

import (
	"fmt"
	// "time"
)

func main() {
	// Basic switch statement example:
	
	// i := 9
	// switch i {
	// case 1:
	// 	fmt.Println("it's one")
	// case 2: 
	// 	fmt.Println("it's two")
	// default :
	// 	fmt.Println("other")
	// }

	/* Explanation:
	   This is a standard `switch` statement in Go.
	   - `i := 9`: We initialize the variable `i` with the value 9.
	   - `switch i`: The switch compares the value of `i` to the cases.
	   - `case 1, case 2`: If `i` equals 1 or 2, it prints respective outputs.
	   - `default`: Executes when none of the above cases match, acting like an `else`.
	*/

	// Switch with multiple conditions:
	
	// switch time.Now().Weekday(){
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("it's the weekend")

	// default:
	// 	fmt.Println("it's a weekday")	
	// }

	/* Explanation:
	   This switch statement checks if the current day is a weekend or a weekday.
	   - `time.Now().Weekday()`: Retrieves the current day of the week.
	   - `case time.Saturday, time.Sunday`: Matches if today is either Saturday or Sunday.
	   - `default`: Catches all other days, indicating they are weekdays.
	*/

	// Type switch example

	// Define an anonymous function `whoami` to determine the type of a variable:
	whoami := func(i interface{}) {
		switch i.(type) { // The `type` keyword allows checking the dynamic type of `i`
		case int: 
			fmt.Println("it's an integer")
		case string:
			fmt.Println("it's a string")
		default:
			fmt.Println("other type")
		}
	}

	// Call the function `whoami` with an integer argument:
	whoami(50)
}

/* Explanation:
   This is a "type switch" which checks the type of the variable `i`.
   - `i interface{}`: Takes an `interface{}`, allowing `i` to hold any data type.
   - `switch i.(type)`: Evaluates the dynamic type of `i`.
   - `case int, case string`: Executes specific cases based on the data type of `i`.
   - `default`: Executes if `i` does not match any specified type.
   
   Here, `whoami(50)` will print "it's an integer" because 50 is of type `int`.
*/
