package pointersTypes

type Person struct {
	Name string
	Age  int
}

func AddAge(user *Person) {
	user.Age++
}

func IncArray(data []int) {
	for i := 0; i < len(data); i++ {
		data[i]++
	}
}

func IncNumber(x *int) {
	*x++
}
