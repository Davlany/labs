package animalfactory_test

import (
	animalfactory "lab2/AnimalFactory"
	"reflect"
	"testing"
)

func TestCat_Say(t *testing.T) {
	testingCat := new(animalfactory.Cat)
	testSay := testingCat.Say()
	if testSay != animalfactory.CatSay {
		t.Errorf("Test fail, expected %s, got %s", animalfactory.CatSay, testSay)
	}
}

func TestDog_Say(t *testing.T) {
	testingDog := new(animalfactory.Dog)
	testSay := testingDog.Say()
	if testSay != animalfactory.CatSay {
		t.Errorf("Test fail, expected %s, got %s", animalfactory.DogSay, testSay)
	}

}

func TestAnimalFactoryDog(t *testing.T) {
	testingData := "dog"
	expected := animalfactory.Animal1
	animal, err := animalfactory.AnimalFactory(testingData)
	if err != nil {
		t.Error(err)
	}
	if reflect.TypeOf(animal).Name() != expected {
		t.Errorf("Test fail, expected %s, got %s", expected, animal)
	}

}

func TestAnimalFactoryCat(t *testing.T) {
	testingData := "cat"
	expected := animalfactory.Animal2
	animal, err := animalfactory.AnimalFactory(testingData)
	if err != nil {
		t.Error(err)
	}
	if reflect.TypeOf(animal).Name() != expected {
		t.Errorf("Test fail, expected %s, got %s", expected, animal)
	}

}
