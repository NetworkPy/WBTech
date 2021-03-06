package main

import "fmt"

// Что выводит данная программа и почему?
// При вызове update копируется указатель p. Получается, что p внутри функции update указывает
// на ту же область, что и p из мейна. Если вы просто присвоите значение по этому указателю (через *p = 2),
// то поменяете именно содержимое этой области памяти.
// Если вы сделаете p = &a, то получите новый указатель на другую область памяти (где хранится a)

func update(p *int) {
	// b := 2
	*p = 23
}

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	update(p)
	fmt.Println(*p)
}
