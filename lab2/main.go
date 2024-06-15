package main

import "fmt"

type Animal interface {
	Say()
}

type Dog struct {
}

type Cat struct {
}

func (d Dog) Say() {
	fmt.Println("Gaw")
}

func (c Cat) Say() {
	fmt.Println("Meow")
}

func AnimalFactory(anType string) Animal {
	if anType == "dog" {
		return Dog{}
	}
	if anType == "cat" {
		return Cat{}
	}
	return nil
}

func main() {
	cat := AnimalFactory("cat")
	cat.Say()
}
