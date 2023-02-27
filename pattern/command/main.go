package main

import "fmt"

type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

type command interface {
	execute()
}

type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

type device interface {
	on()
	off()
}

type auto struct {
	isRunning bool
}

func (a *auto) on() {
	a.isRunning = true
	fmt.Println("Машина заведена")
}

func (a *auto) off() {
	a.isRunning = false
	fmt.Println("Машина заглушена")
}

func main() {
	auto := &auto{}
	onCommand := &onCommand{
		device: auto,
	}
	offCommand := &offCommand{
		device: auto,
	}
	onButton := &button{
		command: onCommand,
	}
	onButton.press()
	offButton := &button{
		command: offCommand,
	}
	offButton.press()
}
