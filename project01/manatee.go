package main

import "fmt"

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

func (m *Manatee) diplay() {
	fmt.Println("Number: %s \n Sex: %s \n Age: %s Size: %s", m.number, m.sex,
		m.age, m.size)
}

/* For sorting, each list needs to be sorted based on age of the manatee. From that point
a comparison across arrays needs to be done to ensure that the manatee is of the correct
size in order to be in that specific order. If not find another manatee of the same age of
the same sex and compare the size of that. If possible to swap do so, if not, impossible.
May need to find permutation in order to find the correct order of manatees to ensure that
it is in the correct order.
*/
