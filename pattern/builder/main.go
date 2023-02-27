package main

import "fmt"

type Auto struct {
	Engine       string
	Transmission string
	Hp           int
}

type BuildAutoI interface {
	Engine(val string) BuildAutoI
	Transmission(val string) BuildAutoI
	Hp(val int) BuildAutoI
	Build() Auto
}

type buildAuto struct {
	engine       string
	transmission string
	hp           int
}

func NewBuildAuto() buildAuto {
	return buildAuto{}
}

func (b buildAuto) Engine(val string) buildAuto {
	b.engine = val
	return b
}

func (b buildAuto) Transmission(val string) buildAuto {
	b.transmission = val
	return b
}

func (b buildAuto) Hp(val int) buildAuto {
	b.hp = val
	return b
}

func (b buildAuto) Bulid() Auto {
	return Auto{
		Engine:       b.engine,
		Transmission: b.transmission,
		Hp:           b.hp,
	}
}

func main() {
	builderAuto := NewBuildAuto()
	auto := builderAuto.Engine("2jz").Transmission("robot").Hp(249)
	fmt.Println(auto)
}
