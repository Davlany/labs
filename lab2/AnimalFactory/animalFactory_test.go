package animalfactory_test

import (
	animalfactory "lab2/AnimalFactory"
	"reflect"
	"testing"
)

func TestCat_Say(t *testing.T) {
	testingCat := new(animalfactory.Cat)
	if testingCat.Say() == "Meow" {
		t.Log("Test complete")
	} else {
		t.Errorf("Test fail, expected Meow, got %s", testingCat.Say())
	}
}

func TestDog_Say(t *testing.T) {
	testingDog := new(animalfactory.Dog)
	if testingDog.Say() == "Gaw" {
		t.Log("Test complete")
	} else {
		t.Errorf("Test fail, expected Gaw, got %s", testingDog.Say())
	}

}

func TestAnimalFactoryDog(t *testing.T) {
	testingData := []string{"dog"}
	expected := []string{"Dog"}
	for i, data := range testingData {
		animal, err := animalfactory.AnimalFactory(data)
		if err != nil {
			t.Error(err)
		}
		if reflect.TypeOf(animal).Name() == expected[i] {
			t.Log("Test complete")
		} else {
			t.Errorf("Test fail, expected %s, got %s", expected[i], animal)
		}
	}
}

func TestAnimalFactoryCat(t *testing.T) {
	testingData := []string{"cat"}
	expected := []string{"Cat"}
	for i, data := range testingData {
		animal, err := animalfactory.AnimalFactory(data)
		if err != nil {
			t.Error(err)
		}
		if reflect.TypeOf(animal).Name() == expected[i] {
			t.Log("Test complete")
		} else {
			t.Errorf("Test fail, expected %s, got %s", expected[i], animal)
		}
	}
}

func TestAnimalFactoryErr(t *testing.T) {
	testingData := []string{"cat", "dog"}
	for _, data := range testingData {
		_, err := animalfactory.AnimalFactory(data)
		if err != nil {
			t.Error(err)
		} else {
			t.Log("Test complete")
		}

	}
}
