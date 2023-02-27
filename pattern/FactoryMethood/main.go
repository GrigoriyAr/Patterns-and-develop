package main

import "fmt"

type iProduct interface { // Определяет все методы которые должны быть у товара в интернет магазине
	setName(n string)
	setWeight(s uint)
	getName() string
	getWeight() uint
}

type beef struct {
	meat
}

func (b beef) setVolume(s uint) {
	panic("implement me")
}

func (b beef) getVolume() uint {
	panic("implement me")
}

func newBeef() *beef {
	return &beef{
		meat: meat{
			name:   "Beef",
			weight: 15,
		},
	}
}
func getProduct(pp string) (iProduct, error) {
	if pp == "beef" {
		return newBeef(), nil
	}
	if pp == "pork" {
		return newPork(), nil
	}
	return nil, fmt.Errorf("Wrong type")
}

type pork struct {
	meat
}

func (p pork) setVolume(s uint) {
	panic("implement me")
}

func (p pork) getVolume() uint {
	panic("implement me")
}

func newPork() *pork {
	return &pork{
		meat: meat{
			name:   "Pork",
			weight: 30,
		},
	}
}

type meat struct {
	name   string
	weight uint
}

func main() {

	bf, _ := getProduct("beef")
	pk, _ := getProduct("pork")

	fmt.Println(bf)
	fmt.Println(pk)
}

func (m *meat) setName(n string) {
	m.name = n
}

func (m *meat) setWeight(s uint) {
	m.weight = s
}

func (m *meat) getName() string {
	return m.name
}

func (m *meat) getWeight() uint {
	return m.weight
}
