package pointersTypes_test

import (
	pointersTypes "labs/lab3/pointers"
	"testing"
)

func TestIncArray(t *testing.T) {
	testTable := []struct {
		arr      []int
		expected []int
	}{
		{[]int{105, 22, 13, -241}, []int{106, 23, 14, -240}},
		{[]int{67, 33, 256, -567}, []int{68, 34, 257, -566}},
		{[]int{45, 16, 678, -1156}, []int{46, 17, 679, -1155}},
		{[]int{80, 11, 1325, -214}, []int{81, 12, 1326, -213}},
	}

	for i, tData := range testTable {
		pointersTypes.IncArray(tData.arr)
		errFlag := false
		for j := 0; j < len(tData.expected); j++ {
			if tData.arr[j] != tData.expected[j] {
				errFlag = true
			}
		}
		if errFlag {
			t.Errorf("TEST № %d EXPECTED %v \t RESULT \t %s%v%s \t❌", i, tData.expected, "\033[31m", tData.arr, "\033[0m")
		} else {
			t.Logf("TEST № %d EXPECTED %v \t RESULT \t %s%v%s \t✅", i, tData.expected, "\033[32m", tData.arr, "\033[0m")
		}
		errFlag = false
	}
}

func TestIncNumber(t *testing.T) {
	testTable := []struct {
		num      int
		expected int
	}{
		{10, 11},
		{32, 33},
		{13, 14},
		{12, 13},
	}

	for i, tData := range testTable {
		pointersTypes.IncNumber(&tData.num)

		if tData.num != tData.expected {
			t.Errorf("TEST № %d EXPECTED %d \t RESULT \t %s%d%s \t❌", i, tData.expected, "\033[31m", tData.num, "\033[0m")
		} else {
			t.Logf("TEST № %d EXPECTED %d \t RESULT \t %s%d%s \t✅", i, tData.expected, "\033[32m", tData.num, "\033[0m")
		}
	}
}

func TestAddAge(t *testing.T) {
	testTable := []struct {
		person   pointersTypes.Person
		expected pointersTypes.Person
	}{
		{person: pointersTypes.Person{Name: "David", Age: 19}, expected: pointersTypes.Person{Name: "David", Age: 20}},
		{person: pointersTypes.Person{Name: "Jack", Age: 21}, expected: pointersTypes.Person{Name: "Jack", Age: 22}},
		{person: pointersTypes.Person{Name: "George", Age: 13}, expected: pointersTypes.Person{Name: "George", Age: 14}},
		{person: pointersTypes.Person{Name: "Floyd", Age: 87}, expected: pointersTypes.Person{Name: "Floyd", Age: 88}},
	}

	for i, tData := range testTable {
		pointersTypes.AddAge(&tData.person)

		if tData.person.Age != tData.expected.Age {
			t.Errorf("TEST № %d EXPECTED %v \t RESULT \t %s%v%s \t❌", i, tData.expected, "\033[31m", tData.person, "\033[0m")
		} else {
			t.Logf("TEST № %d EXPECTED %v \t RESULT \t %s%v%s \t✅", i, tData.expected, "\033[32m", tData.person, "\033[0m")
		}
	}
}
