package animal_test

import (
	animalfactory "lab2/AnimalFactory"
	"testing"
)

func AnimalTest(t *testing.T) {
	testData := []string{"cat", "cat", "dog", "cat", "dog", "dog", "dog"}
	expectedData := []string{"Meow", "Meow", "Gaw", "Meow", "Gaw", "Gaw", "Gaw"}
	for i, v := range testData {
		animal, err := animalfactory.AnimalFactory(v)
		if err != nil {
			s := animal.Say()
			if s == expectedData[i] {
				t.Log("Passed")
			} else {
				t.Errorf("Error, expected value %s, get %s", expectedData[i], v)
			}
		} else {
			t.Errorf(err.Error())
		}
	}

}
