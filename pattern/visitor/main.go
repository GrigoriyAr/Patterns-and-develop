package main

import "fmt"

type video interface {
	WriteAMessage()
}

type youtube struct {
	video
}

func newYoutube() *youtube {
	return &youtube{russian{}}
}

func (y *youtube) setLanguage(video1 video) {
	y.video = video1
}

func (y *youtube) nextLanguage() {
	switch y.video.(type) {
	case english:
		y.video = english{}
	case espaniol:
		y.video = espaniol{}
	case russian:
		y.video = russian{}
	}
}

type english struct {
}

func (en english) WriteAMessage() {
	fmt.Println("Включены английские субтитры")
}

type russian struct {
}

func (ru russian) WriteAMessage() {
	fmt.Println("Включены российские субтитры")
}

type espaniol struct {
}

func (fr espaniol) WriteAMessage() {
	fmt.Println("Включены испанские субтитры")
}

func main() {
	y := newYoutube()
	y.setLanguage(russian{})
	y.WriteAMessage()
}
