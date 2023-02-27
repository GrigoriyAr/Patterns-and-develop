/*
Утилита cut
Реализовать утилиту аналог консольной команды cut (man cut). Утилита должна принимать
строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.
Реализовать поддержку утилитой следующих ключей:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
*/


package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var f int
	var d string
	var s bool

	flag.IntVar(&f, "f", 0, "fields")
	flag.StringVar(&d, "d", "\t", "delimiter")
	flag.BoolVar(&s, "s", false, "separated")

	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(`После завершения ввода введите "exit"`)
	var words [][]string
	for {
		ok := scanner.Scan()
		if !ok {
			log.Fatal(errors.New("Ошибка чтения"))
		}
		line := scanner.Text()
		if line == "exit" {
			break
		}
		if !(s && !strings.Contains(line, d)) {
			words = append(words, strings.Split(line, d))
		}
	}
	if f < 0 {
		log.Fatal(errors.New("Указан некорректныый флаг"))
	}
	if f != 0 {
		var columns []string
		for _, s := range words {
			columns = append(columns, s[f])
		}
		fmt.Println(columns)
	} else {
		for _, s := range words {
			for _, word := range s {
				fmt.Print(word + d)
			}
			fmt.Println("")
		}

	}
}
