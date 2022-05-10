package models

type Person struct {
	Id   int32
	Name string
}

//
func NewPerson(id int32, name string) *Person {
	return &Person{Id: id, Name: name}
}

func (p *Person) SetName(name string) {
	p.Name = name
}
