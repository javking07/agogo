package main

import "fmt"

type person struct {
	fname   string
	lname   string
	favFood []string
}

func (p person) walk() string {
	return fmt.Sprintln(p.fname, "is walking")
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

func (t truck) transportationDevice() string {
	return fmt.Sprintln("This truck hauls")
}

type sedan struct {
	vehicle
	luxury bool
}

func (s sedan) transportationDevice() string {
	return fmt.Sprintln("This sedan rides smooth")
}

type transportation interface {
	transportationDevice() string
}

func report(t transportation) {
	fmt.Println(t.transportationDevice())
}

type gator int

func (g gator) greet() {
	fmt.Println("Sup, I'm a gator", g)
}

type flamingo bool

func (f flamingo) greet() {
	fmt.Println("Hello, I am pink and beautiful and wonderful.")
}

type swampCreature interface {
	greet()
}

func bayou(s swampCreature) {
	s.greet()
}

func main() {
	//testSlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	//testMap := map[string]int{"one": 1, "two": 2, "three": 3}

	// for i, v := range testMap {
	// 	fmt.Println(i, v)
	// }

	// p1 := person{
	// 	"javin",
	// 	"forrester",
	// 	[]string{"pancakes", "bacon", "chicken"},
	// }

	// fmt.Println(p1.fname, p1.lname)
	// for _, v := range p1.favFood {
	// 	fmt.Println(v)
	// }

	// fmt.Println(p1.walk())

	// t1 := truck{
	// 	vehicle:   vehicle{4, "blue"},
	// 	fourWheel: true,
	// }

	// s1 := sedan{
	// 	vehicle: vehicle{4, "red"},
	// 	luxury:  true,
	// }

	// // fmt.Println(t1, "\n", s1)
	// // fmt.Println(t1.color, "\n", s1.luxury)

	// report(t1)
	// report(s1)

	// var g1 gator = 5
	// var f1 flamingo
	//var number = 3
	//number = int(g1)

	// fmt.Printf("g1 is %v and of type %T", g1, g1)
	// fmt.Printf("\nnumber is %v and of type %T", number, number)

	// g1.greet()

	// bayou(g1)
	// bayou(f1)
	s := "i'm sorry dave i can't do that"
	fmt.Println(s)
	fmt.Println([]byte(s))
	fmt.Println(string([]byte(s)))
	fmt.Println(s[:15])
	fmt.Println(s[15:22])
	fmt.Println(s[17:])

	for _, v := range s {
		fmt.Println(string(v))
	}
}
