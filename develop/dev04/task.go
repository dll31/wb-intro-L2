package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
// from https://gist.github.com/johnwesonga/6301924?permalink_comment_id=4704656#gistcomment-4704656
func UniqueSliceElements[T comparable](inputSlice []T) []T {
	uniqueSlice := make([]T, 0, len(inputSlice))
	seen := make(map[T]bool, len(inputSlice))
	for _, element := range inputSlice {
		if !seen[element] {
			uniqueSlice = append(uniqueSlice, element)
			seen[element] = true
		}
	}
	return uniqueSlice
}

func AnagramSearch(inputStrings *[]string) map[string][]string {
	anagrams := make(map[string][]int)

	// make unique every word in dictionary
	uniqueStrings := UniqueSliceElements(*inputStrings)

	for i, word := range uniqueStrings {
		sortedWordByletters := []rune(word)
		slices.SortFunc(sortedWordByletters, func(a, b rune) int {
			return cmp.Compare(a, b)
		})

		// sortedWordByletters is equal for two words from one anagram set

		anagrams[string(sortedWordByletters)] = append(anagrams[string(sortedWordByletters)], i)

	}

	realAnagrams := make(map[string][]string)

	for _, value := range anagrams {
		if len(value) > 1 {
			realKey := uniqueStrings[value[0]]
			realAnagrams[realKey] = make([]string, 0, len(value)-1)
			for i := 1; i < len(value); i++ {
				realAnagrams[realKey] = append(realAnagrams[realKey], uniqueStrings[value[i]])
			}
		}
	}

	// sort values in realAnagrams
	for key := range realAnagrams {
		slices.Sort(realAnagrams[key])
	}

	return realAnagrams
}

func main() {
	file, err := os.Open("dictionary.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0, 1024)

	for scanner.Scan() {
		lines = append(lines, strings.ToLower(scanner.Text()))
	}

	fmt.Printf("%v\n", AnagramSearch(&lines))
}
