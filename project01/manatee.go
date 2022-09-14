package main

import (
	"fmt"
	"sort"
)

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

func processManatees() {
	femaleArray = sortByAge(femaleArray) // Sort array based on age of manatees
	maleArray = sortByAge(maleArray)     // Sort array based on age of manatees
	fmt.Println("Sort based on age.")
	fmt.Println(femaleArray)
	fmt.Println(maleArray)
	organizeBasedOnSize() // Compute the output and arrange manatees accordingly
}

/* For sorting, each list needs to be sorted based on age of the manatee. From that point
a comparison across arrays needs to be done to ensure that the manatee is of the correct
size in order to be in that specific order. If not find another manatee of the same age of
the same sex and compare the size of that. If possible to swap do so, if not, impossible.
May need to find permutation in order to find the correct order of manatees to ensure that
it is in the correct order.
*/

func sortByAge(arr []Manatee) []Manatee {
	sort.SliceStable(arr[:], func(i, j int) bool {
		return arr[i].age < arr[j].age
	})
	return arr
}

// The following function interface lets us sort the array of manatees by age then size.
// after the Manatees are sorted you just have to check if there is a larger manatee in front of a smaller one.
/*
type byAgeFirst []Manatee

func (manatees byAgeFirst) Len() int      { return len(manatees) }
func (manatees byAgeFirst) Swap(i, j int) { manatees[i], manatees[j] = manatees[j], manatees[i] }
func (manatees byAgeFirst) Less(i, j int) bool {
	if manatees[i].age != manatees[j].age {
		return manatees[i].age < manatees[j].age
	}
	return manatees[i].size < manatees[j].size
}
*/


func organizeBasedOnSize() {
	possible := true // Assume that the manatees can be outputted in an order than fits requirements
	// var femaleOutput []Manatee
	// var maleOutput []Manatee

	for i := 0; i < numberInEachRow; i++ {
		if !isValidOutput() {
			fmt.Println(i)
		} else {
			break
		}
	}

	if !isValidOutput() {
		possible = false
	}

	// for i := 0; i < numberInEachRow; i++ {
	// 	if femaleArray[i].size < maleArray[i].size {
	// 		for _, value := range maleArray {
	// 			if maleArray[i].age == value.age && maleArray[i].size > value.size {

	// 			}
	// 		}
	// 	}
	// }

	if !possible { // If output is not acheiveable
		fmt.Println("Impossible")
	} else {
		output() // Print output in correct format
	}
}

func isValidOutput() bool {
	for i := 0; i < numberInEachRow; i++ {
		if femaleArray[i].size < maleArray[i].size {
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
