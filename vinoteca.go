package main

import (
	"fmt"
)

// Vinoteca ...
type Vinoteca struct {
	vinos map[int]*Vino
}

// Vino ...
type Vino struct {
	ID     int
	Name   string
	tipo   string
	año    int
	precio int
}

// NewVinoteca ...
func NewVinoteca() Vinoteca {
	vinos := make(map[int]*Vino)
	return Vinoteca{
		vinos,
	}
}

// Add ...
func (p Vinoteca) Add(v Vino) {
	p.vinos[v.ID] = &v
}

// Print ...
func (p Vinoteca) Print() {
	for _, g := range p.vinos {
		fmt.Printf("ID: %s\n", g.ID)
		fmt.Printf("NOMBRE: %s\n", g.Name)
		fmt.Printf("TIPO: %s\n", g.tipo)
		fmt.Printf("AÑO: %d\n\n", g.año)
		fmt.Printf("PRECIO: %d\n\n", g.precio)
	}
}
