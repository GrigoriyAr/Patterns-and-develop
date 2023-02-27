/*
Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:
- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*
Так же требуется поддерживать функционал fork/exec-команд
Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).
*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).
*/


package main

import (
	"bufio"
	"fmt"
	"github.com/mitchellh/go-ps"
	"os"
	"strconv"
	"strings"
)

func commands(stringCommand string) {
	switch strings.Split(stringCommand, " ")[0] {

	case "cd":
		chDirCommand(stringCommand)
	case "pwd":
		pwdCommand()
	case "echo":
		echoCommand(stringCommand)
	case "kill":
		killPsCommand(stringCommand)
	case "ps":
		psCommand()
	case `\quit`:
		exitCommand()
	default:
		fmt.Println("Такой команды нет")
	}
}

func chDirCommand(stringCommand string) {
	err := os.Chdir(strings.Replace(stringCommand, "cd", "", 1))
	if err != nil {
		fmt.Println(err)
	}
}

func pwdCommand() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir)
}

func echoCommand(stringCommand string) {
	fmt.Println(strings.Replace(stringCommand, "echo", "", 1))
}

func killPsCommand(stringCommand string) {
	pid, err := strconv.Atoi(strings.Replace(stringCommand, "kill", "", 1))
	if err != nil {
		fmt.Println(err)
	}
	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println(err)
	}
	proc.Kill()
}

func psCommand() {
	sliceProc, _ := ps.Processes()
	for _, proc := range sliceProc {
		fmt.Printf("Process name: %v process id: %v\n", proc.Executable(), proc.Pid())
	}
}

func exitCommand() {
	fmt.Println("Exit")
	os.Exit(0)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		commands(scanner.Text())
	}
}
