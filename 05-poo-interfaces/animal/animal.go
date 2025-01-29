package animal

import "fmt"

type Animal interface {
	Sonido()
}

type Perro struct {
	Nombre string
}

func (p *Perro) Sonido() {
	fmt.Println(p.Nombre + "hace goff goff")
}

type Gato struct {
	Nombre string
}

func (p *Gato) Sonido() {
	fmt.Println(p.Nombre + "hace mew mew")
}

func HacerSonido(animal Animal) {
	animal.Sonido()
}
