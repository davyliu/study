package main

import "fmt"

type Bird struct {
	name string
}

func (b *Bird) SetName(name string)  {
	b.name = name
}

func (b *Bird) GetName() string {
	return b.name
}

func (b *Bird) FlyLow()  {
	fmt.Println("I am flying lower...")
}

func (b *Bird) FlyHigh() {
	fmt.Println("I am flying higher...")
}

type IIntro interface {
	SetName(name string)
	GetName() string
}

type IFly interface {
	FlyLow()
	FlyHigh()
}

func main() {
	b := new(Bird)
	var fly IFly = b
	var intro IIntro = b
	intro.SetName("Poly")
	fmt.Println("My name is", intro.GetName())
	fly.FlyLow()
	fly.FlyHigh()
}
