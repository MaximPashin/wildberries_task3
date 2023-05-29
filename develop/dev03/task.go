package main

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

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	column     *uint
	numeric    *bool
	reverse    *bool
	unique     *bool
	month      *bool
	ignoreTail *bool
	check      *bool
)

func main() {
	column = flag.Uint("k", 0, "column")
	numeric = flag.Bool("n", false, "sort as number")
	reverse = flag.Bool("r", false, "sort in reverse order")
	unique = flag.Bool("u", false, "sort without duplicates")
	month = flag.Bool("M", false, "sort by Month name")
	ignoreTail = flag.Bool("b", false, "ignore whitespaces in tail")
	check = flag.Bool("c", false, "check array sorting")
	file := flag.String("file", "", "file for sorting")
	if *numeric && *month {
		log.Fatal("Numeric and month sort at same time")
	}
	if *file == "" {
		log.Fatal("No file passed")
	}
	bdata, err := os.ReadFile(*file)
	if err != nil {
		log.Fatal(err.Error())
	}
	data := []string{}
	lines := strings.Split(string(bdata), "\n")
	if *column > 0 {
		for _, line := range lines {
			data = append(data, strings.Split(line, " ")[*column])
		}
	} else {
		data = lines
	}
	if *unique {
		temp := []string{}
		set := make(map[string]any)
		for _, line := range data {
			_, ok := set[line]
			if !ok {
				set[line] = struct{}{}
				temp = append(temp, line)
			}
		}
		data = temp
	}
	var order []int
	if *numeric {
		order, err = numSort(lines)
	} else {
		if *month {
			order, err = monthSort(lines)
		} else {
			order, err = strSort(lines)
		}
	}
	if err != nil {
		log.Fatal(err.Error())
	}
	if *check {
		for i, v := range order {
			if i != v {
				fmt.Println("Lines not sorted")
				return
			}
		}
		fmt.Println("Lines sorted")
		return
	}
	if *reverse {
		temp := []int{}
		for i := len(order) - 1; i >= 0; i++ {
			temp = append(temp, order[i])
		}
		order = temp
	}
	result := []string{}
	for _, v := range order {
		result = append(result, data[v])
	}
	fmt.Println(strings.Join(result, "\n"))
}

func numSort(data []string) ([]int, error) {
	orderedData := []struct {
		val       int
		initOrder int
	}{}
	for i, line := range data {
		numb, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		orderedData = append(orderedData, struct {
			val       int
			initOrder int
		}{val: numb, initOrder: i})
	}
	sort.Slice(orderedData, func(i, j int) bool {
		return orderedData[i].val < orderedData[j].val
	})
	var result []int
	for _, v := range orderedData {
		result = append(result, v.initOrder)
	}
	return result, nil
}

func monthSort(data []string) ([]int, error) {
	orderedData := []struct {
		val       time.Time
		initOrder int
	}{}
	for i, line := range data {
		m, err := time.Parse("Jan", line)
		if err != nil {
			return nil, err
		}
		orderedData = append(orderedData, struct {
			val       time.Time
			initOrder int
		}{val: m, initOrder: i})
	}
	sort.Slice(orderedData, func(i, j int) bool {
		return orderedData[i].val.Before(orderedData[j].val)
	})
	var result []int
	for _, v := range orderedData {
		result = append(result, v.initOrder)
	}
	return result, nil
}

func strSort(data []string) ([]int, error) {
	orderedData := []struct {
		val       string
		initOrder int
	}{}
	for i, line := range data {
		orderedData = append(orderedData, struct {
			val       string
			initOrder int
		}{val: line, initOrder: i})
	}
	sort.Slice(orderedData, func(i, j int) bool {
		return orderedData[i].val < orderedData[j].val
	})
	var result []int
	for _, v := range orderedData {
		result = append(result, v.initOrder)
	}
	return result, nil
}
