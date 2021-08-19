package main

func main() {
	subaru := &Subaru{18}
	bmw := &Bmw{19}
	audi := &Audi{20}

	printer := &Printer{}

	subaru.accept(printer)
	bmw.accept(printer)
	audi.accept(printer)
}
