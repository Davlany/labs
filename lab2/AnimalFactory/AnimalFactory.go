package animalfactory

import (
	"errors"
)

type Animal interface {
	Say() string
}

type Dogich struct {
}

type Cat struct {
}

func (d Dogich) Say() string {
	return "Gaw"
}

func (c Cat) Say() string {
	return "Meow"
}

func AnimalFactory(anType string) (Animal, error) {
	if anType == "dog" {
		return Dogich{}, nil
	}
	if anType == "cat" {
		return Cat{}, nil
	}
	return nil, errors.New("wrong type")
}
