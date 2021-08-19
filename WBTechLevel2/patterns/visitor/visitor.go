package main

type visitor interface {
	visitForSubaru(*Subaru)
	visitForBmw(*Bmw)
	visitForAudi(*Audi)
}
