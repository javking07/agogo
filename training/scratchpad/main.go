/*
Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of TYPE []string which stores their favorite things. Store three records in your map. Print out all of the values, along with their index position in the slice.


		`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`

		`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`

		`no_dr`, `Being evil`, `Ice cream`, `Sunsets`
*/

package main

import "fmt"

func main() {

	x1 := []string{`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`}
	x2 := []string{`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`}
	x3 := []string{`no_dr`, `Being evil`, `Ice cream`, `Sunsets`}

	xx := [][]string{x1, x2, x3}

	faves := map[string][]string{}

	for _, j := range xx {
		faves[j[0]] = j[1:]
	}

	delete(faves, `moneypenny_miss`)

	for k, v := range faves {
		fmt.Println(k)
		fmt.Println("-----")
		for i, v := range v {
			fmt.Println("\t", i, ":", v)
		}
	}

}
