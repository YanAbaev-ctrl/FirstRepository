package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rimtoarab = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6,
	"VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение: ")
	receiving_values, _ := reader.ReadString('\n')
	receiving_values = strings.TrimSpace(receiving_values)

	if receiving_values == "exit" {
		fmt.Println("Вы вышли из калькулятора.")
		os.Exit(0)
	}

	answer, err := calculate(receiving_values)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Ответ:", answer)
}

func calculate(enter string) (string, error) {
	parts := strings.Fields(enter)
	if len(parts) != 3 {
		panic("Вы не правильно ввели выражение!")
	}

	first_num, operator, second_num := parts[0], parts[1], parts[2]

	rimfirstnum := checkrimnumeral(first_num)
	rimsecondnum := checkrimnumeral(second_num)

	if rimfirstnum != rimsecondnum {
		panic("Не смешивайте две системы исчисления")
	}

	var a, b int
	var err error
	if rimfirstnum {
		a, b, err = parseRomanOperands(first_num, second_num)
	} else {
		a, b, err = parseOperands(first_num, second_num)
	}
	if err != nil {
		return "", err
	}

	var result int
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			panic("на ноль делить нельзя!")
		}
		result = a / b
	default:
		panic("не понятный оператор!")
	}

	if rimfirstnum {
		if result <= 0 {
			panic("римские числа должны быть только положительными")
		}
		return toRoman(result), nil
	}

	return strconv.Itoa(result), nil
}

func parseOperands(first_num, second_num string) (int, int, error) {
	a, err := strconv.Atoi(first_num)
	if err != nil {
		return 0, 0, err
	}
	b, err := strconv.Atoi(second_num)
	if err != nil {
		return 0, 0, err
	}

	if a < 1 || a > 10 || b < 1 || b > 10 {
		panic("не соответствует диапазону [1:10]")
	}

	return a, b, nil
}

func parseRomanOperands(first_num, second_num string) (int, int, error) {
	a, check := rimtoarab[first_num]
	if !check {
		panic("неправильно набрано римское число")
	}
	b, check := rimtoarab[second_num]
	if !check {
		panic("неправильно набрано римское число")
	}

	return a, b, nil
}

func checkrimnumeral(enter string) bool {
	_, check := rimtoarab[enter]
	return check
}

func toRoman(num int) string {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syb := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	roman := ""
	for i := 0; num > 0; i++ {
		for num >= val[i] {
			num -= val[i]
			roman += syb[i]
		}
	}
	return roman
}
