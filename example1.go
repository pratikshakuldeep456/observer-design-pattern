package main

import "fmt"

// observer pattern

type Observerr interface {
	update(string)
	getId() string
}
type Customer struct {
	email string
}

func (c *Customer) update(itemName string) {
	fmt.Println(" sending email to customer", c.email, itemName)
}
func (c *Customer) getId() string {
	return c.email
}

// subject
type Subjectt interface {
	register()
	deregister()
	notifyAll()
}

type Item struct {
	//list of subscribers
	list        []Observerr
	name        string
	isAvailable bool
}

func newItem(name string) *Item {
	return &Item{name: name}
}

func (i *Item) register(o Observerr) {
	i.list = append(i.list, o)

}
func (i *Item) deregister(o Observerr) {
	i.list = removeFromSlicee(i.list, o)

}
func (i *Item) notifyAll() {

	for _, observer := range i.list {
		observer.update(i.name)
	}
}

func (i *Item) updateAvailibilty() {
	fmt.Println("item availablce", i.name)
	i.isAvailable = true
	i.notifyAll()

}

func removeFromSlicee(list []Observerr, o Observerr) []Observerr {

	l := len(list)
	for i, observer := range list {
		if observer.getId() == o.getId() {
			list[l-1], list[i] = list[i], list[l-1]

			return list[:l-1]
		}
	}

	return list
}

func mainn() {

	item := &Item{
		name: "Shirt",
	}

	c1 := &Customer{email: " abc@fmail.com"}
	c2 := &Customer{
		email: " abcd@fmail.com"}

	item.register(c1)
	item.register(c2)
	item.updateAvailibilty()

}
