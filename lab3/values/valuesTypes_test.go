package valuesTypes_test

import (
	valuesTypes "lab3/values"
	"testing"
)

func TestIncNumber(t *testing.T) {
	testTable := []struct {
		num      int
		expected int
	}{
		{10, 10},
		{32, 32},
		{13, 13},
		{12, 12},
	}

	for i, tData := range testTable {
		valuesTypes.IncNumber(tData.num)

		if tData.num != tData.expected {
			t.Errorf("TEST № %d EXPECTED %d \t RESULT \t %s%d%s \t❌", i, tData.expected, "\033[31m", tData.num, "\033[0m")
		} else {
			t.Logf("TEST № %d EXPECTED %d \t RESULT \t %s%d%s \t✅", i, tData.expected, "\033[32m", tData.num, "\033[0m")
		}
	}
}

func TestAddAge(t *testing.T) {
	testTable := []struct {
		person   valuesTypes.Person
		expected valuesTypes.Person
	}{
		{person: valuesTypes.Person{Name: "David", Age: 19}, expected: valuesTypes.Person{Name: "David", Age: 19}},
		{person: valuesTypes.Person{Name: "Jack", Age: 21}, expected: valuesTypes.Person{Name: "Jack", Age: 21}},
		{person: valuesTypes.Person{Name: "George", Age: 13}, expected: valuesTypes.Person{Name: "George", Age: 13}},
		{person: valuesTypes.Person{Name: "Floyd", Age: 87}, expected: valuesTypes.Person{Name: "Floyd", Age: 87}},
	}

	for i, tData := range testTable {
		valuesTypes.AddAge(tData.person)

		if tData.person.Age != tData.expected.Age {
			t.Errorf("TEST № %d EXPECTED %v \t RESULT \t %s%v%s \t❌", i, tData.expected, "\033[31m", tData.person, "\033[0m")
		} else {
			t.Logf("TEST № %d EXPECTED %v \t RESULT \t %s%v%s \t✅", i, tData.expected, "\033[32m", tData.person, "\033[0m")
		}
	}
}
