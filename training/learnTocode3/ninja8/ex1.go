package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByLast []user

func (f ByLast) Len() int           { return len(f) }
func (f ByLast) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f ByLast) Less(i, j int) bool { return (f[i].Last < f[j].Last) }

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return (a[i].Age < a[j].Age) }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	userprint := func(x []user) {
		for _, j := range x {
			fmt.Println("--------")
			fmt.Println(j.Last)
			fmt.Println(j.Age)
			sort.Strings(j.Sayings)
			for i, k := range j.Sayings {
				fmt.Println("Quote ", i)
				fmt.Println(k)
			}
		}

	}
	fmt.Println("----------------------Normal-------------------------")
	userprint(users)
	fmt.Println("----------------------By Last-------------------------")
	sort.Sort(ByLast(users))
	// your code goes here
	userprint(users)
	fmt.Println("----------------------By Age-------------------------")
	sort.Sort(ByAge(users))
	// your code goes here
	userprint(users)
}
