/*
Утилита sort
Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание
и основные параметры): на входе подается файл из несортированными строками,
 на выходе — файл с отсортированными.
Реализовать поддержку утилитой следующих ключей:
-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки
Дополнительно
Реализовать поддержку утилитой следующих ключей:
-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учетом суффиксов
*/


package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var column int

func initFlags() {
	var (
		number    bool
		reverse   bool
		unique    bool
		month     bool
		backSpace bool
		check     bool
		numberH   bool
	)
	flag.IntVar(&column, "k", 0, "указание колонки для сортировки")
	flag.BoolVar(&number, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&reverse, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&unique, "u", false, "не выводить повторяющиеся строки")
	flag.BoolVar(&month, "M", false, "сортировка по названию месяца")
	flag.BoolVar(&backSpace, "b", false, "игнорировать хвостовые пробелы")
	flag.BoolVar(&check, "c", false, "провверять отсортированы-ли данные")
	flag.BoolVar(&numberH, "h", false, "сортировать по числовому знаению с учетом суффиксов")
	flag.Parse()
}

func main() {
	initFlags()
	result, err := readFile("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func readFile(filename string) (result []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if column > 0 {
			strArr := strings.Split(scanner.Text(), " ")
			result = append(result, strArr[column])
		} else {
			result = append(result, scanner.Text())
		}
	}
	sort.Strings(result)
	return
}
