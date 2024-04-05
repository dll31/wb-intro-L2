package main

import (
	"errors"
	"fmt"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Unpack(s *string) (string, error) {
	runes := []rune(*s)

	var escape bool = false
	var sequence []rune

	for i, r := range runes {
		if r == '\\' && !escape {
			escape = true
			continue
		}

		if unicode.IsDigit(r) && i == 0 {
			return "", errors.New("bad string")
		} else if unicode.IsDigit(r) && unicode.IsDigit(runes[i-1]) && runes[i-2] != '\\' {
			return "", errors.New("bad string")
		} else if unicode.IsDigit(r) && !escape {
			c := int(r - '0')

			for j := 0; j < c-1; j++ {
				sequence = append(sequence, runes[i-1])
			}
		} else {
			sequence = append(sequence, r)
			escape = false
		}
	}

	return string(sequence), nil
}

func main() {
	input := []string{"a4bc2d5e", "abcd", "45", "", "qwe\\4\\5", "qwe\\45", "qwe\\\\5"}

	for _, s := range input {
		out, err := Unpack(&s)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(out)
		}
	}
}
