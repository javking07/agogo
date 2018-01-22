/*
Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:

first name
last name
favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.
*/

package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	faves []string
}

func show(p person) {

	fmt.Println(p.first)
	fmt.Println(p.last)
	for _, v := range p.faves {
		fmt.Printf("\t%v\n", v)
	}

}
func main() {

	p1 := person{
		first: "Javin",
		last:  "Forrester",
		faves: []string{"chocolate", "rum and raisin"},
	}

	p2 := person{
		first: "Kristie",
		last:  "Qiu",
		faves: []string{"Butter", "pikachu"},
	}

	personMap := make(map[string]person)

	personMap[p1.last] = p1
	personMap[p2.last] = p2

	show(personMap["Forrester"])
	show(personMap["Qiu"])

}
