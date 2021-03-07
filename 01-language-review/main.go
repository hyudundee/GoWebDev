package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good morning."`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "Shaken, not stirred."`)
}

func main() {
	// x := 7
	// fmt.Printf("%T", x)
	p1 := person{
		"Miss",
		"Moneypenny",
	}
	// p1.speak()

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	// sa1.speak()
	// sa1.person.speak()
	saySomething(p1)
	saySomething(sa1)
}
