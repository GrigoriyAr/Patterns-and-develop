package main

import "fmt"

type auto struct {
}

func (a auto) onPetrolPump() {
	fmt.Println("Бензонасос работает")
}

func (a auto) onLight() {
	fmt.Println("Фары включены")
}

func (a auto) onEngine() {
	fmt.Println("Двигатель заведен")
}

func (a auto) onMirror() {
	fmt.Println("Зеркала выдвинуты")
}

func (a auto) onBelt() {
	fmt.Println("Ремень выдвинут")
}

type facadeAuto struct {
	auto *auto
}

func (f *facadeAuto) onAuto() {
	f.auto.onPetrolPump()
	f.auto.onLight()
	f.auto.onEngine()
	f.auto.onMirror()
	f.auto.onBelt()
}

func main() {
	autom := &auto{}
	facade := &facadeAuto{autom}
	facade.onAuto()
}
