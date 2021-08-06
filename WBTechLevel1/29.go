package main

import (
	"fmt"
	"math/big"
)

// Написать программу, которая перемножает, делит, складывает, вычитает 2 числовых переменных a,b, значение которые > 2^20.

func main() {
	value1 := new(big.Int)
	value1.SetString("44000000000000000000", 10)

	value2 := new(big.Int)
	value2.SetString("24000000000000000000", 10)

	fmt.Println(MyDiv(value1, value2))
	fmt.Println(MyAdd(value1, value2))
	fmt.Println(MyMul(value1, value2))
	fmt.Println(MySub(value1, value2))
}

// "/"
func MyDiv(a *big.Int, b *big.Int) *big.Int {
	value := new(big.Int)
	value.Div(a, b)
	return value
}

// "+"
func MyAdd(a *big.Int, b *big.Int) *big.Int {
	value := new(big.Int)
	value.Add(a, b)
	return value
}

// "*"
func MyMul(a *big.Int, b *big.Int) *big.Int {
	value := new(big.Int)
	value.Mul(a, b)
	return value
}

// "-"
func MySub(a *big.Int, b *big.Int) *big.Int {
	value := new(big.Int)
	value.Sub(a, b)
	return value
}
