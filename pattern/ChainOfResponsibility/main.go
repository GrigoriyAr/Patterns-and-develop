package main

import "fmt"

type part interface { // Интерфейсу должны удовлетворять обработчики
	execute(*clothes)
	setNext(part)
}

type clothes struct {
	isCreated bool
	isWashes  bool
	isPayment bool
}

type dryCleaning struct {
	next part
}

func newDryCleaning() *dryCleaning {
	return &dryCleaning{}
}

func (d *dryCleaning) execute(order *clothes) {
	fmt.Println("Заказ создам")
	order.isCreated = true
	d.next.execute(order)
}

func (d *dryCleaning) setNext(department part) {
	d.next = department
}

type dryClean struct {
	next part
}

func (d *dryClean) execute(order *clothes) {
	if !order.isCreated {
		return
	}
	fmt.Println("Заказ прошел чистку")
	order.isWashes = true
	d.next.execute(order)
}

func (d *dryClean) setNext(department part) {
	d.next = department
}

type Cashier struct {
	next part
}

func (c *Cashier) execute(order *clothes) {
	if !order.isWashes {
		return
	}
	fmt.Println("Заказ оплачен")
	order.isPayment = true
}

func (c *Cashier) setNext(next part) {
	c.next = next
}

func main() {

	cashier := &Cashier{}

	store := &dryClean{}
	store.setNext(cashier)

	storage := &dryCleaning{}
	storage.setNext(store)

	order := &clothes{}
	storage.execute(order)
}
