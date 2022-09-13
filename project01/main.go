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

// Global variable for count of manatees in a row
var numberInEachRow int

func main() {
	takeInput() // Call take input function
}

/*
Take input from standard input stream and store the data in some container. List? Heap?
*/

func takeInput() {
	const inputRows = 4
	// Delete later:
	fmt.Print("Enter number of manatees per row: ")
	fmt.Scan(&numberInEachRow) // Take input for number of manatees

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
		ageArray = trim(ageArray)                 // Trim off excess \n if necessary

		if len(ageArray) != numberInEachRow || !isValidInput(ageArray) { // Check to see if input meets the requirements
			ageArray = retakeInput()
		}

		ageArr := cleanArray(ageArray) // Clean the data and store in an array of type int

		counter++

		// Read the second portion of the input for the sex, this portion corresponds to the size of the manatee in the row.

		sizeString, err2 := reader.ReadString('\n') // Read in data from Stdin
		if err2 != nil {
			panic(err2)
		}
		sizeArray := strings.Split(sizeString, " ")
		sizeArray = trim(sizeArray) // Trim off excess \n if necessary

		if len(sizeArray) != numberInEachRow || !isValidInput(sizeArray) { // Check to see if input meets requirements
			sizeArray = retakeInput()
		}

		sizeArr := cleanArray(sizeArray) // Convert array to type int

		counter++

		/*
			Add each variable from arrays to its corresponding variable in manatee object. Arrays are specified for
			variable. Number and sex can be determined based on the number of iterations in the array as well
			as the overall loop.
		*/

		for i := 0; i < numberInEachRow; i++ {
			var m Manatee // Initialize Manatee object
			m.number = i + 1
			m.sex = sex
			m.age = ageArr[i]
			m.size = sizeArr[i]

			if sex == "Female" {
				femaleArray = append(femaleArray, m)
			} else {
				maleArray = append(maleArray, m)
			}
		}
	}
	fmt.Println(femaleArray) // Delete later, testing purposes.
	fmt.Println(maleArray)
}

/*
Check whether input is valid, by checking to see if all elements in array are of type int. Will fail in the event that there is
excess whitespace or a letter in the array.
*/

func isValidInput(arr []string) bool {
	for _, value := range arr { // Iterate through array
		_, err := strconv.Atoi(value) // Convert value to int
		if err != nil {               // If value can not be converted then the input is invalid
			return false
		}
	}
	return true
}

func cleanArray(stringArr []string) []int {
	var intArr []int
	for i := 0; i < len(stringArr); i++ {
		intVar, err1 := strconv.Atoi(stringArr[i]) // Convert string to int
		var stringAge string
		if err1 != nil { // If the int throws an error, process that error accordingly and clean the value to be assigned
			stringAge = strings.TrimSuffix(stringArr[i], "\n") // Remove suffix from the string.
			intVar, err1 := strconv.Atoi(stringAge)            // Convert the string to an int
			if err1 != nil {                                   // Ensure that it does not throw an expection error
				panic(err1)
			}
			intArr = append(intArr, intVar)
		} else { // If no problems rise populate variable in object
			intArr = append(intArr, intVar)
		}
	}
	return intArr
}

/*
If input is invlaid, retake input will ensure that the input being given is proper. Will not allow to continue with
execution until the given input meets the requirements for execution. Checking to see if data is the same size as well
as the data only contains integers.
*/

func retakeInput() []string {
	reader := bufio.NewReader(os.Stdin)
	var stringArr []string
	for len(stringArr) != numberInEachRow || !isValidInput(stringArr) {
		fmt.Println("Invalid Input. Please try again.")
		ageString, err1 := reader.ReadString('\n')
		if err1 != nil {
			panic(err1)
		}
		stringArr = strings.Split(ageString, " ")
		if len(stringArr) == numberInEachRow {
			stringArr = trim(stringArr)
		}
	}
	return stringArr
}

/*
	Trim function takes the array and cleans up the elements to ensure that there is no newline operand as a suffix.
*/

func trim(arr []string) []string {
	var temp []string
	for i := 0; i < len(arr); i++ {
		_, err := strconv.Atoi(arr[i])
		if err != nil {
			element := strings.Trim(arr[i], "\n")
			temp = append(temp, element)
		} else {
			temp = append(temp, arr[i])
		}
	}
	arr = temp
	return arr
}
