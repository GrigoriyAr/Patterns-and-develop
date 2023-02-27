/*
Задача на распаковку
Создать Go-функцию, осуществляющую примитивную распаковку строки,
содержащую повторяющиеся символы/руны, например:
"a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (некорректная строка)
"" => ""
Дополнительно
Реализовать поддержку escape-последовательностей.
Например:
qwe\4\5 => qwe45 (*)
qwe\45 => qwe44444 (*)
qwe\\5 => qwe\\\\\ (*)
В случае если была передана некорректная строка,
функция должна возвращать ошибку. Написать unit-тесты.
*/


package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func Unpack(strings string) (string, error) {
	runeString := []rune(strings)
	stringEdited := ""
	var k int

	if len(runeString) == 0 {
		return "", nil
	}
	if unicode.IsDigit(runeString[0]) {
		return "некорректная строка", nil
	}

	for i := 1; i < len(runeString); i++ {
		if unicode.IsLetter(runeString[i-1]) && unicode.IsDigit(runeString[i]) {
			s, _ := strconv.Atoi(string(runeString[i]))
			for j := 0; j < s; j++ {
				stringEdited += string(runeString[i-1])
				k += j + 1
			}
		} else if unicode.IsLetter(runeString[i-1]) == true {
			stringEdited += string(runeString[i-1])
			k += 1
		}
	}

	if unicode.IsLetter(runeString[len(runeString)-1]) == true {
		stringEdited += string(runeString[len(runeString)-1])
	}

	return stringEdited, nil
}

func main() {
	firstString := "a4bc2d5e"
	fmt.Println(Unpack(firstString))

	secondString := "abcd"
	fmt.Println(Unpack(secondString))

	thirdString := "45"
	fmt.Println(Unpack(thirdString))

	fourthString := ""
	fmt.Println(Unpack(fourthString))
}
