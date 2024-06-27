package animalfactory_test

import (
	animalfactory "lab2/AnimalFactory"
	"reflect"
	"testing"
)

func TestCat_Say(t *testing.T) {
	testingCat := new(animalfactory.Cat)
	testSay := testingCat.Say()
	if testSay != "Meow" {
		t.Errorf("Test fail, expected %s, got %s", "Meow", testSay)
	}
}

func TestDog_Say(t *testing.T) {
	testingDog := new(animalfactory.Dogich)
	testSay := testingDog.Say()
	if testSay != "Gaw" {
		t.Errorf("Test fail, expected %s, got %s", "Gaw", testSay)
	}

}

func TestAnimalFactoryDog(t *testing.T) {
	animal, _ := animalfactory.AnimalFactory("dog")
	animalType := reflect.TypeOf(animal).Name()
	dogType := reflect.TypeOf(animalfactory.Dogich{}).Name()

	if animalType != dogType {
		t.Errorf("Test failed, expected %s got: %s", dogType, animalType)
	}

}

func TestAnimalFactoryCat(t *testing.T) {
	animal, _ := animalfactory.AnimalFactory("cat")
	animalType := reflect.TypeOf(animal).Name()
	catType := reflect.TypeOf(animalfactory.Cat{}).Name()

	if animalType != catType {
		t.Errorf("Test failed, expected %s got: %s", catType, animalType)
	}
}

func TestErrors(t *testing.T) {
	testValues := []string{"Bird", "dock", "", "papich"}
	for _, v := range testValues {
		_, err := animalfactory.AnimalFactory(v)
		if err == nil {
			t.Errorf("Test failed, expected error: wrong type, but error is nil")
		}
	}
}
