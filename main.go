package main

func main() {
	var max = 10
	p := NewVinoteca()
	for i := 0; i < max; i++ {
		p.Add(Vino{i, "uxmal", "malbec", 2009, 180})
	}

	p.Print()
}
