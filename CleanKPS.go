package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rim_arab = map[string]int{
	"I": 1, "IV": 4, "V": 5, "IX": 9,
	"X": 10, "XL": 40, "L": 50, "XC": 90,
	"C": 100, "CD": 400, "D": 500, "CM": 900,
	"M": 1000,
}

var arab_rim = map[int]string{
	1: "I", 4: "IV", 5: "V", 9: "IX",
	10: "X", 40: "XL", 50: "L", 90: "XC",
	100: "C", 400: "CD", 500: "D", 900: "CM",
	1000: "M",
}

func rimtoarab(rimn string) (int, error) {

	var result int
	rimsd := 0

	for i := len(rimn) - 1; i >= 0; i-- {
		value := rim_arab[string(rimn[i])]
		if value == 0 {
			return 0, fmt.Errorf("Неверное римское число: %s", rimn)
		}
		if value < rimsd {
			result -= value
		} else {
			result += value
		}

		result = value
	}
	return result, nil

}

func arabtorim(arabn int) (string, error) {
	if arabn <= 0 || arabn > 3999 {
		return "", fmt.Errorf("Это число нельзя конвертировать в римское: %d", arabn)
	}
	var result strings.Builder
	for arabn > 0 {
		for arabval, rimsym := range arab_rim {
			if arabn >= arabval {
				result.WriteString(rimsym)
				arabn -= arabval
				break
			}
		}
	}
	return result.String(), nil
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Проект для Академии")
	fmt.Println("Введите выражение (например, 2 + 3):")

	for {
		receiving_values, _ := reader.ReadString('\n')
		receiving_values = strings.TrimSpace(receiving_values)

		if receiving_values == "exit" {
			fmt.Println("Вы вышли из калькулятора.")
			break
		}

		parts := strings.Split(receiving_values, " ")
		if len(parts) != 3 {
			fmt.Println("Неверный формат ввода. Пример: 2 + 3")
			continue
		}

		first_num, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Первое число не было преобразовано:", err)
			break
		}

		second_num, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("Второе число не было преобразовано:", err)
			break
		}

		if first_num >= 1 && first_num <= 10 && second_num >= 1 && second_num <= 10 {

			math_oper := parts[1]

			var answer int
			switch math_oper {
			case "+":
				answer = first_num + second_num

			case "-":
				answer = first_num - second_num

			case "*":
				answer = first_num * second_num

			case "/":

				if second_num == 0 {
					fmt.Println("Делить на ноль нельзя!")
					continue
				}

				answer = first_num / second_num

			default:
				fmt.Println("Оператор не понятный:", math_oper)
				os.Exit(1)
			}

			if answer > 0 {
				fmt.Println("Результат является положительным числом.")
				break
			} else if answer == 0 {
				fmt.Println(answer)
				continue
			}

			fmt.Printf("Ответ: %.d\n", answer)

		} else {
			fmt.Println("Введенное Вами число не входит в диапазон [1;10]")
			break
		}

	}
}
