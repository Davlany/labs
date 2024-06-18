package animalfactory

import (
	"errors"
)

type Animal interface {
	Say() string
}

type Dog struct {
}

type Cat struct {
}

func (d Dog) Say() string {
	return "Gaw"
}

func (c Cat) Say() string {
	return "Meow"
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
