package models

type personManager struct {
	persons []*Person
}

func NewPersonManager() personManager {
	return personManager{persons: make([]*Person, 0)}
}

func (p *personManager) AddPerson(person *Person) {
	p.persons = append(p.persons, person)
}

func (p *personManager) Length() int {
	return len(p.persons)
}

func (p *personManager) GetPersons() []*Person {
	return p.persons
}

func (p *personManager) GetPerson(id int32) *Person {

	for _, v := range p.persons {
		if v.Id == id {
			return v
		}
	}

	return nil
}
