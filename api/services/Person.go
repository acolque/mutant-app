package services

type Person struct {
	name string
}

func (p Person) GetName() string {
	return "Arnold"
}
