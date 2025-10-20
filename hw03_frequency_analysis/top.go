package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	// Требование теста: на пустой строке вернуть пустой срез, а не nil.
	if len(strings.TrimSpace(text)) == 0 {
		return []string{}
	}

	// 1) Разделяем только по пробельным символам.
	tokens := strings.Fields(text)
	if len(tokens) == 0 {
		return []string{}
	}

	// 2) Считаем частоты.
	freq := make(map[string]int, len(tokens))
	for _, t := range tokens {
		freq[t]++
	}

	// 3) Перекладываем в пары и сортируем:
	//    - по частоте по убыванию
	//    - при равенстве — лексикографически по возрастанию
	type pair struct {
		word string
		cnt  int
	}
	arr := make([]pair, 0, len(freq))
	for w, c := range freq {
		arr = append(arr, pair{word: w, cnt: c})
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i].cnt != arr[j].cnt {
			return arr[i].cnt > arr[j].cnt
		}
		return arr[i].word < arr[j].word
	})

	// 4) Берём топ-10 слов
	n := 10
	if len(arr) < n {
		n = len(arr)
	}
	out := make([]string, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, arr[i].word)
	}
	return out
}
