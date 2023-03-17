package main

import (
	"os"

	unixtable "github.com/frantjc/go-encoding-unixtable"
)

type Person struct {
	ID   string `unixtable:"-"`
	Name string `unixtable:"Full Name"`
	Age  int
}

func GetPerson(id string) *Person {
	return &Person{
		ID:   id,
		Name: "Obi-Wan Kenobi",
		Age:  35,
	}
}

func GetPeople() []Person {
	return []Person{
		{
			ID:   "1",
			Name: "Obi-Wan Kenobi",
			Age:  35,
		},
		{
			ID:   "2",
			Name: "General Grievous",
			Age:  50,
		},
	}
}

func main() {
	var p any

	if len(os.Args) > 1 {
		p = GetPerson(os.Args[1])
	} else {
		p = GetPeople()
	}

	if err := unixtable.NewEncoder(os.Stdout).Encode(p); err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
