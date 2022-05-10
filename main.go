package main

import (
	"fmt"

	"github.com/ramses2099/datastructureingo/models"
)

type Test struct {
	Id int
}

func NewTest(id int) *Test {
	return &Test{Id: id}
}

var (
	AllTest []*Test
)

func printPerson(p *models.Person) {
	fmt.Printf(" id %v name %v \n", p.Id, p.Name)
}

//
func main() {

	// s := models.Student{Id: 25, Name: "juan perez"}

	// fmt.Println("name :", s.Name)

	// changed(&s)

	// fmt.Println("name :", s.Name)

	// t := NewTest(12)

	// AllTest = append(AllTest, t)

	// fmt.Println("befored Id ", t.Id)
	// t.Id = 50

	// fmt.Println("after Id ", t.Id)

	// fmt.Println("Legnth allTest ", len(AllTest))

	// for _, v := range AllTest {
	// 	fmt.Println("Element ", v.Id)
	// 	v.Id = 40
	// }

	// fmt.Println("after Id ", t.Id)

	p1 := models.NewPerson(1, "juan perez")
	p2 := models.NewPerson(2, "juana perez")
	p3 := models.NewPerson(3, "juan carlos canela")

	pm := models.NewPersonManager()
	pm.AddPerson(p1)
	pm.AddPerson(p2)
	pm.AddPerson(p3)

	for _, v := range pm.GetPersons() {
		printPerson(v)
		v.Name = v.Name + " santorio"
	}
	//
	fmt.Println("after changing")

	for _, v := range pm.GetPersons() {
		printPerson(v)
	}

	pu := pm.GetPerson(1)
	fmt.Println("search ", pu.Name)
	pu.Name = "juan perez"
	fmt.Println("change search ", pu.Name)

}
