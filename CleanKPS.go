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

var arabtorim = []string{
	"", "I", "II", "III", "IV", "V", "VI",
	"VII", "VIII", "IX", "X", "XI", "XII",
	"XIII", "XIV", "XV", "XVI", "XVII", "XVIII",
	"XIX", "XX", "XXI", "XXII", "XXIII", "XXIV",
	"XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX",
	"XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI",
	"XXXVII", "XXXVIII", "XXXIX", "XL", "XLI", "XLII",
	"XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII",
	"XLIX", "L", "LI", "LII", "LIII", "LIV",
	"LV", "LVI", "LVII", "LVIII", "LIX", "LX",
	"LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI",
	"LXVII", "LXVIII", "LXIX", "LXX", "LXXI", "LXXII",
	"LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII",
	"LXXIX", "LXXX", "LXXXI", "LXXXII", "LXXXIII", "LXXXIV",
	"LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC",
	"XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI",
	"XCVII", "XCVIII", "XCIX", "C",
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
		return "", fmt.Errorf("Вы не правильно ввели выражение")
	}

	first_num, operator, second_num := parts[0], parts[1], parts[2]

	rimfirstnum := checkrimnumeral(first_num)
	rimsecondnum := checkrimnumeral(second_num)

	if rimfirstnum != rimsecondnum {
		return "", fmt.Errorf("Смешивать две системы исчисления не стоит!")
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
			return "", fmt.Errorf("На ноль делить нельзя!")
		}
		result = a / b
	default:
		return "", fmt.Errorf("Неизвестное действие")
	}

	if rimfirstnum {
		if result <= 0 {
			return "", fmt.Errorf("Римские числа должны быть только положительными")
		}
		return arabtorim[result], nil
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
		return 0, 0, fmt.Errorf("Не соответствует диапазону [1:10]")
	}

	return a, b, nil
}

func parseRomanOperands(first_num, second_num string) (int, int, error) {
	a, check := rimtoarab[first_num]
	if !check {
		return 0, 0, fmt.Errorf("Не правильно набрано римское число: %s", first_num)
	}
	b, check := rimtoarab[second_num]
	if !check {
		return 0, 0, fmt.Errorf("Не правильно набрано римское число: %s", second_num)
	}

	return a, b, nil
}

func checkrimnumeral(enter string) bool {
	_, check := rimtoarab[enter]
	return check
}
