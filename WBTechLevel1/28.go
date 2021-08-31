package main

import "fmt"

// Реализовать паттерн адаптер на любом примере.
// Адаптер выступает прослойкой между двумя объектами, превращая вызовы одного в вызовы понятные другому.

func main() {
	cli := &client{}
	usb := &USB{}
	typec := &typeC{}
	cli.insertUSB(usb)

	adapter := &typeCadapter{
		tc: typec,
	}
	cli.insertUSB(adapter)
}

type connector interface {
	InsertUSB()
}

type client struct {
}

func (c *client) insertUSB(con connector) {
	con.InsertUSB()
}

type USB struct {
}

func (u *USB) InsertUSB() {
	fmt.Println("Insert USB")
}

type typeC struct {
}

func (t *typeC) InsertTypeC() {
	fmt.Println("Insert TypeC")
}

type typeCadapter struct {
	tc *typeC
}

func (t *typeCadapter) InsertUSB() {
	t.tc.InsertTypeC()
}
