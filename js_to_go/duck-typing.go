package main

type Duck interface {
	Quacks()
}

type Animal struct {
	Name string
}

func (a Animal) Quacks() {
	println(a.Name, "quacks")
}

func Scream(duck Duck) {
	duck.Quacks()
}

func main() {
	d := Animal{"Duck"}
	d.Quacks()

	Scream(d)
}
