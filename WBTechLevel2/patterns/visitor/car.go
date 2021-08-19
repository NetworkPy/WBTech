package main

type Car interface {
	getType() string
	accept(visitor)
}
