package main

import (
	"errors"
	"strconv"
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

type packedSymb struct {
	symb  rune
	times uint
}

// ErrIncorrectString error throw in case Unpack arguments incorrect form
var ErrIncorrectString = errors.New("Unpack string: Incorrect input string")

// Unpack packed string to normal form
func Unpack(str string) (string, error) {
	// обработка пустой строки
	if len(str) == 0 {
		return str, nil
	}
	// флаг - в эскейп последовательности или нет
	esc := []rune(str)[0] == '\\'
	// флаг - в количестве повторений или нет
	inNumb := false
	// последний встречанный символ
	var lastSymb *packedSymb
	// инициализация переменных
	if !esc {
		if !unicode.IsDigit([]rune(str)[0]) {
			lastSymb = &packedSymb{symb: []rune(str)[0], times: 1}
		} else {
			return "", ErrIncorrectString
		}
	}
	// временный массив отвечающий за запакованную строку
	packedStr := make([]packedSymb, 0)
	// парсинг запакованной строки
	for _, symb := range str[1:] {
		if !esc && unicode.IsDigit(symb) {
			if lastSymb != nil {
				if inNumb {
					lastSymb.times *= 10
					numb, _ := strconv.Atoi(string(symb))
					lastSymb.times += uint(numb)
				} else {
					numb, _ := strconv.Atoi(string(symb))
					lastSymb.times = uint(numb)
				}
			} else {
				return "", ErrIncorrectString
			}
		} else {
			inNumb = false
			if !esc && symb == '\\' {
				esc = true
			} else {
				if lastSymb != nil {
					packedStr = append(packedStr, *lastSymb)
				}
				lastSymb = &packedSymb{symb: symb, times: 1}
				esc = false
			}
		}
	}
	if esc {
		return "", ErrIncorrectString
	}
	if lastSymb != nil {
		packedStr = append(packedStr, *lastSymb)
	}
	// распаковка строки в нормальную форму
	result := []rune{}
	for _, pSymb := range packedStr {
		temp := []rune{}
		for i := uint(0); i < pSymb.times; i++ {
			temp = append(temp, pSymb.symb)
		}
		result = append(result, temp...)
	}
	return string(result), nil
}
