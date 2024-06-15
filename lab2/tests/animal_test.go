package an_test

import "testing"

func AnimalTest(t *testing.T) {
	testData := []string{"cat", "cat", "dog", "cat", "dog", "dog", "dog"}
	expectedData := []string{"Meow","Meow","Gaw","Meow","Gaw","Gaw","Gaw"}
	for i, v := range testData {
		if v != expectedData[i]{
			t.Error("Bad test")
		}
	}

}
