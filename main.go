package main

import (
	"fmt"
)

var LicenseToKill bool

type person struct {
	Name string
	Age  int
}

type doubleZero struct {
	person
	LicenseToKill bool
}
type thripleZero struct {
	doubleZero
	LicenseToKill bool
}

func (p person) Greeting() {
	fmt.Println("I'm just a regular person.")
}

func (dz doubleZero) Greeting() {
	fmt.Println("Miss Moneypenny, so good to see you.")
}
func (dz thripleZero) Greeting() {
	fmt.Println("Miss Moneypenny, I must kill you")
}

func main() {
	p1 := person{
		Name: "Ian Flemming",
		Age:  44,
	}

	p2 := doubleZero{
		person: person{
			Name: "James Bond",
			Age:  23,
		},
		LicenseToKill: true,
	}

	p3 := thripleZero{
		doubleZero : doubleZero{
			person: person{
			Name : "Joshua Agbeku",
			Age : 26,
			},
		},
		LicenseToKill: true,
	}
	p1.Greeting()
	p2.Greeting()
	p2.person.Greeting()

	if ( p3.LicenseToKill == true){//invokes dz thrippleZero function if the condition is true.
	p3.Greeting()
	}
	p3.person.Greeting()

}
