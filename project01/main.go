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
	"strings"
)

// Global variable arrays of type Manatee
var femaleArray Manatee
var maleArray Manatee

func main() {
	// Call take input function
	takeInput()
}

/*
Take input from standard input stream and store the data in some container. List? Heap?
*/

func takeInput() {
	const input_rows = 4
	//Initialize number of manatees in each row
	var numberInEachRow int

	// Delete later:
	fmt.Print("Enter number of manatees per row: ")
	// Take input for number of manatees
	fmt.Scan(&numberInEachRow)

	reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter text: ")

	text, _ := reader.ReadString('\n')

	fmt.Println(text)

	integer := strings.Split(text, " ")

	// fmt.Println(integer)
	// fmt.Println(integer[0])
	// fmt.Println(integer[1])
	// fmt.Println(integer[2])
	// fmt.Println(len(integer))
	valid := true
	i := 0
	// Take input while expecting input to be taken
	for valid && i < inputinput_rows {
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
		integer := strings.Split(text, " ")
		if len(integer) != numberInEachRow {
			// fmt.Println("Input of expected length. Please try again.")
			valid = false
		}
		m = Manatee()
		
		i++
	}




	for i := 0; i < input_rows; i++ {
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
		integer := strings.Split(text, " ")
		if len(integer) != numberInEachRow {
			fmt.Println("Input of expected length. Please try again.")
			while valid
		}
	// 	var row string
	// 	// fmt.Scan(&row)
	// 	te
	// 	integers := strings.Split(row, " ")
	// 	fmt.Println(integers)
	// 	fmt.Println(i)
	}
}
