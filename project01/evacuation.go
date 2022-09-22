/*
 * Author: Spencer Hirsch, shirsch2020@my.fit.edu
 * Author: James Pabisz, jpabisz2020@my.fit.edu
 * Course: CSE 4250, Fall 2022
 * Project: Proj1
 * Implementation: go version gccgo
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
	Intitalize type Manatee. Manatee consists of the number tattooed on the manatee,
	the sex of the manatee, the age of the manatee and the size of the manatee.const
*/

type Manatee struct {
	number int
	sex    string
	age    int
	size   int
}

/*
Take input from standard input stream and store the data in some container. List? Heap?
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
			if numberInEachRow <= 5 || numberInEachRow >= 1 {
				isValid = true
			}
		}
	}

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
	fmt.Println("Inputted values, input order.")
	fmt.Println(femaleArray) // Delete later, testing purposes.
	fmt.Println(maleArray)
}

func cleanArray(stringArr []string) []int {
	var intArr []int
	for i := 0; i < len(stringArr); i++ {
		intVar, err1 := strconv.Atoi(stringArr[i]) // Convert string to int
		// var stringAge string
		if err1 != nil { // If the int throws an error, process that error accordingly and clean the value to be assigned
			panic(err1)
			// stringAge = strings.TrimSuffix(stringArr[i], "\n") // Remove suffix from the string.
			// intVar, err1 := strconv.Atoi(stringAge)            // Convert the string to an int
			// if err1 != nil {                                   // Ensure that it does not throw an expection error
			// 	panic(err1)
			// }
			// intArr = append(intArr, intVar)
		} else { // If no problems rise populate variable in object
			intArr = append(intArr, intVar)
		}
	}
	return intArr
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

func processManatees() {
	femaleArray = sortByAge(femaleArray) // Sort array based on age of manatees
	maleArray = sortByAge(maleArray)     // Sort array based on age of manatees
	fmt.Println("Sort based on age.")
	fmt.Println(femaleArray)
	fmt.Println(maleArray)

	// Find manatees of the same age
	mra := same_ages(maleArray)
	fra := same_ages(femaleArray)
	fmt.Print("female: ")
	fmt.Println(fra)
	fmt.Print("male: ")
	fmt.Println(mra)
	// organizeBasedOnSize() // Compute the output and arrange manatees accordingly
}

func sortByAge(arr []Manatee) []Manatee {
	sort.SliceStable(arr[:], func(i, j int) bool {
		return arr[i].age < arr[j].age
	})
	return arr
}

func contains(val int, arr []int) bool {
	for _, value := range arr {
		if value == val {
			return true
		}
	}
	return false
}

func isValidOutput() bool {
	for i := 0; i < numberInEachRow; i++ {
		if femaleArray[i].size <= maleArray[i].size {
			return false
		}
	}
	return true
}

func output() {
	for _, value := range femaleArray {
		fmt.Print(value.number)
		fmt.Print(" ")
	}
	fmt.Print("\n")
	for _, value := range maleArray {
		fmt.Print(value.number)
		fmt.Print(" ")
	}
	fmt.Print("\n")
}

func organizeBasedOnSize() {
	for i := 0; i < numberInEachRow; i++ {
		if !isValidOutput() {
			fmt.Println(i)
		} else {
			break
		}
	}

	if !isValidOutput() { // If output is not acheiveable
		fmt.Println("impossible")
	} else {
		output() // Print output in correct format
	}
}

func same_ages(arr []Manatee) [][]Manatee {
	var repeatAge []int
	for _, value := range arr {
		age := value.age
		count := 0
		for _, val2 := range arr {
			if age == val2.age {
				count++
			}
		}

		if count >= 1 && !contains(age, repeatAge) {
			repeatAge = append(repeatAge, age)
		}
	}
	fmt.Println("repeat")
	fmt.Println(repeatAge)
	manateeRepeatAge := find_manatee(repeatAge, arr)
	return manateeRepeatAge
}

func find_manatee(repeatAgeIntArr []int, manateeArr []Manatee) [][]Manatee {
	var manateeRepeatAge [][]Manatee
	for _, age := range repeatAgeIntArr {
		var individualAge []Manatee
		for _, manatee := range manateeArr {
			if age == manatee.age {
				individualAge = append(individualAge, manatee)
			}
		}
		manateeRepeatAge = append(manateeRepeatAge, individualAge)
	}
	return manateeRepeatAge
}

func factorial(n int)(result int) {
	if (n > 0) {
		result = n * factorial(n-1)
		return result
	}
	return 1
}


func permutate(sameAge []Manatee) {
     var n = len(sameAge) - 1
     var i, j int
     var numOfPerms = factorial(len(sameAge))
     for c := 1; c < numOfPerms; c++ { // 3! = 6:
            i = n - 1
            j = n
            for sameAge[i].number > sameAge[i+1].number {
                    i--
            }
            for sameAge[j].number < sameAge[i].number {
                    j--
            }
            sameAge[i], sameAge[j] = sameAge[j], sameAge[i]
            j = n
            i += 1
            for i < j {
                    sameAge[i], sameAge[j] = sameAge[j], sameAge[i]
                    i++
                    j--
            }
            fmt.Println(sameAge)
    }
}

func main() {
	takeInput() // Call take input function
	processManatees()
}
