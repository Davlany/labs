package animalfactory

import (
	"errors"
)

const (
	Animal1 = "Dog"
	Animal2 = "Cat"
	DogSay  = "Gaw"
	CatSay  = "Meow"
)

type Animal interface {
	Say() string
}

type Dog struct {
}

type Cat struct {
}

func (d Dog) Say() string {
	return DogSay
}

func (c Cat) Say() string {
	return CatSay
}

func AnimalFactory(anType string) (Animal, error) {
	if anType == "dog" {
		return Dog{}, nil
	}
	if anType == "cat" {
		return Cat{}, nil
	}
	return nil, errors.New("wrong type")
}
