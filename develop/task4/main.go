/*
Поиск анаграмм по словарю
Написать функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.
Требования:
Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
Выходные данные: ссылка на мапу множеств анаграмм
Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
слово из множества.
Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.
*/


package main

import (
	"fmt"
	"sort"
	"strings"
)

func UniqLower(in *[]string) []string {
	res := make([]string, 0, len(*in))
	u := make(map[string]bool)

	for _, i := range *in {
		if !u[i] {
			res = append(res, strings.ToLower(i))
			u[i] = true
		}
	}

	return res
}

func AnagramDict(in *[]string) map[string][]string {
	if len(*in) < 2 {
		return nil
	}

	data := make(map[string][]string)

	uniqIn := UniqLower(in)
	for _, i := range uniqIn {
		sorted := []rune(i)
		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i] < sorted[j]
		})

		word := string(sorted)
		data[word] = append(data[word], i)

	}

	res := make(map[string][]string)
	for _, words := range data {
		if len(words) > 1 {
			sort.Strings(words)
			res[words[0]] = words
		}
	}

	return res
}

func main() {
	input := []string{"тест", "листок", "пятка", "пятак", "тяпка", "листок", "пятка", "слиток", "столик"}
	fmt.Println(AnagramDict(&input))
}
