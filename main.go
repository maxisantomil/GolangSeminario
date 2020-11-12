package main

import (
	"fmt"
)

//printable *interface*

type Printable interface {
	print()
}
type person struct {
	name     string
	lastname string
}

type figure struct {
	name string
	area int
}

//funciones
func (p person) print() {
	fmt.Println("[person]", p.name, p.lastname)
}

// el * apunta al original ... pudiendo hacerle cambios
func (p *person) clear() {
	p.name = " "
}

func (f *figure) print() {
	fmt.Println("[figure]", f.name, f.area)
}

func invokePrint(p Printable) {
	p.print()
}

func main() {
	//crea un puntero de persona, (&) seria como un new
	p := &person{name: "juan", lastname: "santomil"}
	invokePrint(p)

	f := &figure{name: "cube", area: 150}
	invokePrint(f)
}
