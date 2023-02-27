package main

import (
	"fmt"
)

type state interface {
	vioce() string
	vibr() string
}

// Структура основного объекта door. Методы входа, открытия и закрытия структуры вызывают соответствующие методы поля текущего состояния
type phone struct {
	Voice        state
	Vibr         state
	currentState *phone
}

func (p *phone) setState(s state) {
	p.currentState = p
}

// Метод включения звука
func (p *phone) voiceMod() string {
	return "Громкость включена"
}

// Метод выключения звука
func (p *phone) vibrMod() string {
	return "Вибрация включена"
}

// Состояние вкл
type voiceState struct {
	phone *phone
}

// Структура состояния "Закрыто"
type vibrState struct {
	phone *phone
}

// Метод структуры состояния прибавления громкости для выкл звука
func (v *vibrState) turnSound() error {
	return fmt.Errorf("Чтобы прибавить звук включите режим звук ")
}

// Метод структуры состояния вкл звука для для выкл звука
func (v *voiceState) offSound() error {
	fmt.Println("Closing the door")
	v.phone.setState(v.phone.Vibr)
	return nil
}

// Метод структуры состояния состояния выкл звука для убавления громкости
func (p *vibrState) lowSound() error {
	return fmt.Errorf("Чтобы убавить звук включите режим звук ")
}

// Метод структуры для прибавления звука при включенном звуке
func (p *phone) turnSound() string {
	return "Звук прибавлен"
}

func (p *phone) lowSound() string {
	return "Звук убавлен"
}

