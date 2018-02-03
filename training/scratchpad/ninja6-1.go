package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

type shape interface {
	area() float64
}

func info(s shape) {
	x := s.area()
	fmt.Println(x)
}
func giveastring() string {
	return "Hey there"
}

func takeafunc(xs string, f func(xi string) string) string {
	return f(xs)

}
func main() {

	g := func(xi string) string {
		return string(xi[0])
	}
	y := takeafunc("ab", g)

	fmt.Println(y)
}
