package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type sortOptions struct {
	filename  string
	key       int
	numeric   bool
	reverse   bool
	unique    bool
	separator string
}

const startBatchSize = 1024

func main() {
	options := parseOptions()

	file, err := os.Open(options.filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0, startBatchSize)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sort.Slice(lines, func(i, j int) bool {
		a, b := splitLine(lines[i], options.separator), splitLine(lines[j], options.separator)
		valueA, valueB := a[options.key-1], b[options.key-1]

		if options.numeric {
			numA, errA := strconv.Atoi(valueA)
			numB, errB := strconv.Atoi(valueB)

			if errA != nil || errB != nil {
				fmt.Fprintf(os.Stderr, "ошибка: некорректное числовое значение: %s или %s\n", valueA, valueB)
				os.Exit(1)
			}

			return numA < numB
		}

		return valueA < valueB
	})

	if options.reverse {
		reverseSlice(&lines)
	}

	if options.unique {
		uniqueLines := make([]string, 0, startBatchSize)
		prevLine := ""

		for _, line := range lines {
			if line != prevLine {
				uniqueLines = append(uniqueLines, line)
				prevLine = line
			}
		}

		lines = uniqueLines
	}

	printLines(&lines)

	writeOutFile(completeOutFilename(&options), &lines)
}

func parseOptions() sortOptions {
	options := sortOptions{
		filename:  "",
		key:       1,
		numeric:   false,
		reverse:   false,
		unique:    false,
		separator: " ",
	}

	numetricFlag := flag.Bool("n", false, "sort by numeric value")
	reverseFlag := flag.Bool("r", false, "sort in reverse order")
	uniqueFlag := flag.Bool("u", false, "do not output duplicate lines")
	columnFlag := flag.Int("k", 1, "sortable column")

	flag.Parse()

	if *numetricFlag {
		options.numeric = true
	}
	if *reverseFlag {
		options.reverse = true
	}
	if *uniqueFlag {
		options.unique = true
	}
	if *columnFlag > 0 {
		options.key = *columnFlag
	}

	options.filename = flag.Arg(0)

	return options
}

func completeOutFilename(options *sortOptions) string {
	completeFilename := options.filename + "_sorted"
	if options.reverse {
		completeFilename += "_rev"
	}
	if options.unique {
		completeFilename += "_uniq"
	}
	if options.numeric {
		completeFilename += "_num"
	}

	completeFilename += "_col_" + strconv.Itoa(options.key)

	return completeFilename
}

func writeOutFile(filename string, lines *[]string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, line := range *lines {
		_, err := file.Write([]byte(line + "\n"))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func splitLine(line, separator string) []string {
	return strings.Split(line, separator)
}

func reverseSlice(s *[]string) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func printLines(lines *[]string) {
	for _, line := range *lines {
		fmt.Println(line)
	}
}
