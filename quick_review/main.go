package main

import "fmt"

//creating a new type
type person struct {
	fname string
	lname string
}

//creating a new type
type secretAgent struct {
	p   person
	ltk bool
}

//creating an interface
type human interface {
	speak()
}

//creating methods
func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says: "Hello World!"`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.p.fname, sa.p.lname, `says: "I have a license to kill!"`)
}

//function that accepts type human : polymorphism
//if a type has a speak() method, they too are of type human
func saySomething(h human) {
	h.speak()
}

func main() {

	//creating an object of type person
	p1 := person{
		"Jose",
		"Hernandez",
	}

	//creating an object of type secretAgent
	sa1 := secretAgent{
		//have to create a person or pass an already created person
		p: person{
			"Markojudas",
			"Lightbringer",
		},

		ltk: true,
	}

	saySomething(p1)
	saySomething(sa1)
	saySomething(sa1.p) // this works becase sa1.p is type person, which in turn it is also type human!

}
