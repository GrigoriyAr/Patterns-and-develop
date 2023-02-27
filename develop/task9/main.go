/*
Реализовать утилиту wget с возможностью скачивать сайты целиком.
*/


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://garantzakona.ru/" // адрес сайта, который нужно скачать

	resp, err := http.Get(url) // получаем данные с сайта

	if err != nil { // проверяем наличие ошибки
		fmt.Println(err) // выводим ошибку, если таковая имеется
	} else { // если ошибки нет, то:

		defer resp.Body.Close() // удаляем body

		body, err := ioutil.ReadAll(resp.Body) // читаем body

		if err != nil { // проверяем, что body удалось успешно считать
			fmt.Println(err) //выводим ошибку, если таковая имеется
		} else { //выводим body, если body удалось успешно считать

			fmt.Println(string(body))

		}

	}
}
