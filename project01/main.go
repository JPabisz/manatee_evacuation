/*
 * Author: Spencer Hirsch, shirsch2020@my.fit.edu
 * Author: James Pabisz, jpabisz2020@my.fit.edu
 * Course: CSE 4250, Fall 2022
 * Project: project tag, short project name
 * Implementation: compiler version
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Global variable arrays of type Manatee
var femaleArray []Manatee
var maleArray []Manatee

func main() {
	// Call take input function
	takeInput()
}

/*
Take input from standard input stream and store the data in some container. List? Heap?
*/

func takeInput() {
	const inputRows = 4
	//Initialize number of manatees in each row
	var numberInEachRow int

	// Delete later:
	fmt.Print("Enter number of manatees per row: ")
	// Take input for number of manatees
	fmt.Scan(&numberInEachRow)

	reader := bufio.NewReader(os.Stdin)

	counter := 0
	/*
		Take input for as long as input is expected. Use count variable to keep track of the number of interations and
		increment as necessary. This design allows for two inputs to be taken within the same iteration of the loop while incrementing
		the count variable. Purpose is to ensure that the Manatee object is populated correctly for each manatee in the row. As
		well as ensure the sex of the manatees are maintained correctly.
	*/
	for counter < inputRows {
		var sex string
		if counter < inputRows/2 { // While the count variable is less than 2, the sex is Female. Else, male.
			sex = "Female"
		} else {
			sex = "Male"
		}
		ageString, err1 := reader.ReadString('\n') // Read input from Stdin
		if err1 != nil {
			panic(err1)
		}
		ageArray := strings.Split(ageString, " ") // Convert into string array split at the whitespace

		/*
			Check to see if the input is valid. If not allow for the input to be retaken appropriately. Iterate throuhg the loop
			until input is in the correct format to be processed. Primary checks to ensure that the size of the array is the expected
			size given by the user as @var numberInEachRow.
		*/
		if len(ageArray) != numberInEachRow {
			fmt.Println("Invalid input! Try again.")
			var newAgeArray []string                  // Temporary array to hold the newly constructed array given that it is valid
			for len(newAgeArray) != numberInEachRow { // Check to see if the length of the array is correct before proceeding.
				ageString, err1 := reader.ReadString('\n')
				if err1 != nil {
					panic(err1)
				}
				newAgeArray = strings.Split(ageString, " ")
			}
			ageArray = newAgeArray // Assign ageArray as the temporary array to continue with execution
		}
		counter++

		// Read the second portion of the input for the sex, this portion corresponds to the size of the manatee in the row.

		sizeString, err2 := reader.ReadString('\n')
		if err2 != nil {
			panic(err2)
		}
		sizeArray := strings.Split(sizeString, " ")
		if len(sizeArray) != numberInEachRow {
			fmt.Println("Invalid input!")
			var newSizeArray []string                  // Temporary array to hold the newly constructed array given that it is valid
			for len(newSizeArray) != numberInEachRow { // Check to see if the length of the array is correct before proceeding.
				sizeString, err1 := reader.ReadString('\n')
				if err1 != nil {
					panic(err1)
				}
				newSizeArray = strings.Split(sizeString, " ")
			}
			sizeArray = newSizeArray // Assign sizeArray as the temporary array to continue with execution
		}
		counter++

		/*
			This section of code handles the processing of the collected data from the input. By taking the input in the
			arrays and assigning the value at each index to a variable within the Manatee object. Through this it also checks
			to see if whether or not the value contains a suffix of \n and then trims it off if so.
		*/

		for i := 0; i < numberInEachRow; i++ {
			var m Manatee // Initialize Manatee object
			m.number = i + 1
			m.sex = sex
			intAge, err1 := strconv.Atoi(ageArray[i]) // Convert string to int
			var stringAge string
			if err1 != nil { // If the int throws an error, process that error accordingly and clean the value to be assigned
				stringAge = strings.TrimSuffix(ageArray[i], "\n") // Remove suffix from the string.
				intAge, err1 := strconv.Atoi(stringAge)           // Convert the string to an int
				if err1 != nil {                                  // Ensure that it does not throw an expection error
					panic(err1)
				}
				m.age = intAge // Populate the variable in the object
			} else { // If no problems rise populate variable in object
				m.age = intAge
			}
			intSize, err2 := strconv.Atoi(sizeArray[i]) // Take string and convert to int
			var stringSize string
			if err2 != nil { // If error is thrown process the string accordingly, same as lines above
				stringSize = strings.TrimSuffix(sizeArray[i], "\n")
				intSize, err2 := strconv.Atoi(stringSize)
				fmt.Println(stringSize)
				if err2 != nil {
					panic(err2)
				}
				m.size = intSize // Populate variable in object with variable
			} else { // If no errors arise
				m.size = intSize
			}

			if sex == "Female" {
				femaleArray = append(femaleArray, m)
			} else {
				maleArray = append(maleArray, m)
			}
		}
	}
}

// Still need to check for whitespace in arrays as a condition that it is populated properly. Create method that will compute and verify. Return boolean
