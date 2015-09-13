package main

import "fmt"

type speak interface {
	SayHello(name string)
}

type person struct {
	Name string
}

type robot struct {
	Name string
}

func (p *person) SayHello() {
	p.Name = "Jean Matheus Souto"
	fmt.Printf("Hi %s, how are you?\n", p.Name)
}

func (r robot) SayHello() {
	fmt.Printf("Hello i'm your robot, my name is %s\n", r.Name)
}

func main() {
	jean := person{"Jean Matheus"}
	atom := robot{"Atom"}

	jean.SayHello()
	atom.SayHello()

	fmt.Println("----")

	fmt.Println(jean.Name)
	fmt.Println(atom.Name)
}
