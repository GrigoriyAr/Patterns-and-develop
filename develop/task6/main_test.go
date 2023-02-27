package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	// Проверка флагов
	if f != 0 {
		if f < 0 {
			t.Errorf("Указан некорректный флаг")
		} else {
			var columns []string
			for _, s := range words {
				columns = append(columns, s[f])
			}

			if columns != expectedColumns {
				t.Errorf("Ошибка при составлении столбцов")
			}

		} else {

			for _, s := range words {

				for _, word := range s {

					if word+d != expectedWord+d {

						t.Errorf("Ошибка при составлении слов")

					}

				}

				fmt.Println("")

			}

		}
	}
}