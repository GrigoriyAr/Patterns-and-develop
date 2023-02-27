package main

import (
	"reflect"
	"testing"
)

func TestAnagramDict(t *testing.T) {
	test := []string{"тест", "листок", "пятка", "пятак", "тяпка", "листок", "пятка", "слиток", "столик"}
	val_answ := AnagramDict(&test)
	val_answ_true := map[string][]string{"листок": []string{"листок", "слиток", "столик"}, "пятак": []string{"пятак", "пятка", "тяпка"}}

	if !reflect.DeepEqual(val_answ, val_answ_true) {
		t.Fatalf("Test case: %v,\n function returns: %v,\n true answ: %v\n", test, val_answ, val_answ_true)
	}
}
