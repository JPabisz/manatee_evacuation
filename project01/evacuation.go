/*
 * Author: Spencer Hirsch, shirsch2020@my.fit.edu
 * Author: James Pabisz, jpabisz2020@my.fit.edu
 * Course: CSE 4250, Fall 2022
 * Project: Proj1, Manatee Evacuation
 * Implementation: go version go 1.19.1 darwin/arm64
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var femaleArray []Manatee // Global variable arrays of type Manatee
var maleArray []Manatee
var numberInEachRow int // Global variable for count of manatees in a row

/*
	Intitalize type Manatee. Manatee consists of the number tattooed on the
	manatee, the sex of the manatee, the age of the manatee and the size
	of the manatee.const
*/

type Manatee struct {
	number int
	sex    string
	age    int
	size   int
}

/*
Take input from standard input stream and store the data in some container.
*/

func takeInput() {
	const inputRows = 4
	// Delete later:
	fmt.Print("Enter number of manatees per row: ")
	fmt.Scan(&numberInEachRow) // Take input for number of manatees
	if numberInEachRow > 5 || numberInEachRow < 1 {
		isValid := false
		for !isValid {
			fmt.Print("Invalid Input. Enter number of manatees per row: ")
			fmt.Scan(&numberInEachRow) // Take input for number of manatees
			if numberInEachRow <= 5 && numberInEachRow >= 1 {
				isValid = true
			}
		}
	}

	reader := bufio.NewReader(os.Stdin)

	counter := 0

	/*
		Take input for as long as input is expected. Use count variable to
		keep track of the number of interations and increment as necessary.
		This design allows for two inputs to be taken within the same iteration
		of the loop while incrementing the count variable. Purpose is to ensure
		that the Manatee object is populated correctly for each manatee in the
		row. As well as ensure the sex of the manatees are maintained
		correctly.
	*/

	for counter < inputRows {
		var sex string
		// While the count variable is less than 2, the sex is Female.
		if counter < inputRows/2 {
			sex = "Female"
		} else {
			sex = "Male"
		}
		ageString, err1 := reader.ReadString('\n') // Read input from Stdin
		if err1 != nil {
			panic(err1)
		}

		// Convert into string array split at the whitespace
		ageArray := strings.Split(ageString, " ")
		ageArray = trim(ageArray) // Trim off excess \n if necessary
		// Check to see if input meets the requirements
		if len(ageArray) != numberInEachRow || !isValidInput(ageArray) {
			ageArray = retakeInput()
		}

		// Clean the data and store in an array of type int
		ageArr := cleanArray(ageArray)

		counter++
		/*
			Read the second portion of the input for the sex, this portion
			corresponds to the size of the manatee in the row.
		*/
		sizeString, err2 := reader.ReadString('\n') // Read in data from Stdin
		if err2 != nil {
			panic(err2)
		}
		sizeArray := strings.Split(sizeString, " ")
		sizeArray = trim(sizeArray) // Trim off excess \n if necessary

		// Check to see if input meets requirements
		if len(sizeArray) != numberInEachRow || !isValidInput(sizeArray) {
			sizeArray = retakeInput()
		}

		sizeArr := cleanArray(sizeArray) // Convert array to type int

		counter++

		/*
			Add each variable from arrays to its corresponding variable in
			manatee object. Arrays are specified for variable. Number and
			sex can be determined based on the number of iterations in the
			array as well as the overall loop.
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
}

/*
	Function used to clean up the array passed as a parameter. The initial
	array is given as a string. Function then coverts each string into type int
	and appends the newly converted int to a new array. The function then
	returns an array of type int.
*/

func cleanArray(stringArr []string) []int {
	var intArr []int
	for i := 0; i < len(stringArr); i++ {
		intVar, err1 := strconv.Atoi(stringArr[i]) // Convert string to int
		/*
			If the int throws an error, process that error accordingly and
			clean the value to be assigned
		*/
		if err1 != nil {
			panic(err1)
		} else { // If no problems rise populate variable in object
			intArr = append(intArr, intVar)
		}
	}
	return intArr
}

/*
	Trim function takes the array and cleans up the elements to ensure that
	there is no newline operand as a suffix.
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

/*
	Check whether input is valid, by checking to see if all elements in array
	are of type int. Will fail in the event that there is excess whitespace or
	a letter in the array.
*/

func isValidInput(arr []string) bool {
	for _, value := range arr { // Iterate through array
		_, err := strconv.Atoi(value) // Convert value to int

		// If value can not be converted then the input is invalid
		if err != nil {
			return false
		}

		// if val <= 1 || val >= 1,000,000 {
		// 	return false
		// }
	}
	return true
}

/*
	If input is invlaid, retake input will ensure that the input being given is
	proper. Will not allow to continue with execution until the given input
	meets the requirements for execution. Checking to see if data is the same
	size as well as the data only contains integers.
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
	Function sorts each manatee array based on the ages of the manatees. If
	ages are the same the returned array will be of
*/

func sortByAge(arr []Manatee) []Manatee {
	sort.SliceStable(arr[:], func(i, j int) bool {
		return arr[i].age < arr[j].age
	})
	return arr
}

/*
	Contains function checks to see if a value is within an array of
	manatees.
*/

func contains(val int, arr []int) bool {
	for _, value := range arr {
		if value == val {
			return true
		}
	}
	return false
}

/*
Function checks to see if the output is valid. Checks to see if the arrays are
sorted. If so, iterates through each Manatee object and checks the size to
ensure that the back row is not smaller than the front row. If so, return
false.
*/

func isValidOutput(male []Manatee, female []Manatee) bool {
	if !isSorted(male) {
		return false
	} else if !isSorted(female) {
		return false
	} else {
		for i := 0; i < numberInEachRow; i++ {
			if male[i].size >= female[i].size {
				return false
			}
		}
	}
	return true
}

/*
Function hecks to see if the manatees in an array given as a parameter
is in ascending order. Compares current value to the next until it
either fails or passes. Returns a boolean.
*/
func isSorted(arr []Manatee) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i].age > arr[i+1].age {
			return false
		}
	}
	return true
}

/*
	Prints the rows of manatees both male and female. Female in the bacl
	male in the front.
*/

func output(male []Manatee, female []Manatee) {
	for _, value := range female {
		fmt.Print(value.number)
		fmt.Print(" ")
	}
	fmt.Print("\n")
	for _, value := range male {
		fmt.Print(value.number)
		fmt.Print(" ")
	}
	fmt.Print("\n")
}

/*
	Cacualtes the factorial of an integer given as a paramter.
*/

func factorial(n int) (result int) {
	if n > 0 {
		result = n * factorial(n-1)
		return result
	}
	return 1
}

/*
	Function calculates all possible permutations of the male manatees.
	After each permutation check to see if permutation is in a valid
	order such that it fits the criteria. Call the function to calculate
	all permutations of female manatees.
*/

func permutateM() {
	var n = len(maleArray) - 1
	var array []Manatee = maleArray
	var females []Manatee = sortByAge(femaleArray)
	if isValidOutput(array, females) {
		output(array, females)
		return
	}
	var i, j int

	// Calculate number of permutations
	var numOfPerms = factorial(len(maleArray))
	for c := 1; c < numOfPerms; c++ {
		i = n - 1
		j = n

		for array[i].number > array[i+1].number {
			i--
		}
		for array[j].number < array[i].number {
			j--
		}
		array[i], array[j] = array[j], array[i]

		j = n
		i += 1
		for i < j {
			array[i], array[j] = array[j], array[i]
			i++
			j--
		}
		if isValidOutput(array, females) { // Check for valid output
			output(array, females)
			return // terminate function if valid output is reached
		}

	}

	permutateF() // Call permutate function for females

}

/*
Calculate all of the permutations of the female manatee array. After
each permutation check to see if the ordering is valid. If not try
another permutation.
*/
func permutateF() {
	var n = len(femaleArray) - 1
	var array []Manatee = femaleArray
	var males []Manatee = sortByAge(maleArray)
	if isValidOutput(males, array) {
		output(males, array)
		return
	}
	var i, j int
	// Calculate number of permutations using factorial function.
	var numOfPerms = factorial(len(femaleArray))
	for p := 1; p < numOfPerms; p++ {
		i = n - 1
		j = n
		for array[i].number > array[i+1].number {
			i--
		}
		for array[j].number < array[i].number {
			j--
		}
		array[i], array[j] = array[j], array[i]

		j = n
		i += 1
		for i < j {
			array[i], array[j] = array[j], array[i]
			i++
			j--
		}
		// Check to see if output is valid.
		if isValidOutput(males, array) {
			output(males, array)
			return // Terminate the function if output is found
		}

	}

	/*
		If not combination is found to work it is impossible to order them
		in such a way it fits the criteria
	*/
	fmt.Println("impossible")

}

// Main driver for the program. Calls all necessary functions.

func main() {
	takeInput()  // Call take input function
	permutateM() // Call permutate function for males
}
