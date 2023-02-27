package main

import (
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	result, err := readFile("text.txt")
	if err != nil {
		t.Errorf("Ошибка при чтении файла: %v", err)
	}

	expected := []string{"golang intern junior wb"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидаемый результат %v, полученный %v", expected, result)
	}
}
