package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

		first_num, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			fmt.Println("Первое число не было преобразовано:", err)
			continue
		}

		second_num, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			fmt.Println("Второе число не было преобразовано:", err)
			continue
		}

		math_oper := parts[1]

		var answer float64
		switch math_oper {
		case "+":
			answer = first_num + second_num
		case "-":
			answer = first_num - second_num
		case "*":
			answer = first_num * second_num
		case "/":
			if second_num == 0 {
				fmt.Println("Деление на ноль не пройдет!")
				continue
			}
			answer = first_num / second_num
		default:
			fmt.Println("Оператор не понятный:", math_oper)
			continue
		}

		fmt.Printf("Ответ: %.2f\n", answer)
	}
}
