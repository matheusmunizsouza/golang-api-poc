package main

type Person struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	BirthDay string `json:"birthDay"`
}

var people []Person

func findAll() *[]Person {
	if people == nil {
		return &[]Person{}
	}
	return &people
}

func findByName(name string) *Person {
	for _, p := range people {
		if p.Name == name {
			return &p
		}
	}
	return nil
}

func save(person Person) {
	person.Id = len(people) + 1
	people = append(people, person)
}

func update(id int, person Person) {
	for i, p := range people {
		if p.Id == id {
			p.Name = person.Name
			p.LastName = person.LastName
			p.BirthDay = person.BirthDay
			people[i] = p
		}
	}
}

func delete(id int) {
	for i, p := range people {
		if p.Id == id {
			var newPeople []Person
			for i2, p2 := range people {

				if i != i2 {
					newPeople = append(newPeople, p2)
				}
			}
			people = newPeople
		}
	}
}
