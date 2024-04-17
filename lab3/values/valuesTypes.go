package valuesTypes

type Person struct {
	Name string
	Age  int
}

func AddAge(user Person) {
	user.Age++
}

func IncNumber(x int) {
	x++
}
