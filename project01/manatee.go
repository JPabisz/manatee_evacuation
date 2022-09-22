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

func sortBySize(arr [][]Manatee) {
	male := arr[0]
	female := arr[1]

	sort.SliceStable(male[:], func(i, j int) bool {
		return male[i].size < male[j].size
	})
	return male

	sort.SliceStable(female[:], func(i, j int) bool {
		return female[i].size > female[j].size
	})
	return female
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

// The following function interface lets us sort the array of manatees by age then size.
// after the Manatees are sorted you just have to check if there is a larger manatee in front of a smaller one.

// type byAgeFirst []Manatee

// func (manatees byAgeFirst) Len() int      { return len(manatees) }
// func (manatees byAgeFirst) Swap(i, j int) { manatees[i], manatees[j] = manatees[j], manatees[i] }
// func (manatees byAgeFirst) Less(i, j int) bool {
// 	if manatees[i].age != manatees[j].age {
// 		return manatees[i].age < manatees[j].age
// 	}
// 	return manatees[i].size < manatees[j].size
// }

func organizeBasedOnSize() {
	for i := 0; i < numberInEachRow; i++ {
		if !isValidOutput() {
			fmt.Println(i)
		} else {
			break
		}
	}

	if !isValidOutput() { // If output is not acheiveable
		fmt.Println("Impossible")
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

		if count > 1 && !contains(age, repeatAge) {
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

func proc() {
	for i := 0; i < numberInEachRow; i++ {
		if maleArray[i].size > femaleArray[i].size {

		}
	}
}
