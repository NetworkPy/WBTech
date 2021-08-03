package main

import "fmt"

// Реализовать композицию структуры Action от родительской структуры Human.

type Action struct {
	*Human //
	Job    string
}

// Parent
type Human struct {
	Name string
	Age  int
}

func (h *Human) Introduce() {
	fmt.Printf("Hi, I'm %s\n", h.Name)
}

func main() {
	a := &Action{
		Human: &Human{"Jhon Doe", 11},
		Job:   "Worker",
	}
	fmt.Println(a.Human.Name)
	a.Human.Introduce()
}
