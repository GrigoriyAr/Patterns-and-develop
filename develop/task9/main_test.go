package main

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestMain(t *testing.T) {
	url := "https://example.com/"
	resp, err := http.Get(url)

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Error: %s", err)
	} else {
		if string(body) == "" { // проверяем, что body не пустой
			t.Errorf("Body is empty")
		} else { // если body не пустой, то:
			t.Logf("Body is not empty") // выводим сообщение, что body не пустой
		}
	}
}
