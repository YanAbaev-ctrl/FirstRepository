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

		first_num, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Первое число не было преобразовано:", err)
			break
		}

		if first_num < 1 || first_num > 10 {
			fmt.Println("Принимаем только числа от 1 до 10 ключительно!")
			break
		}

		second_num, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("Второе число не было преобразовано:", err)
			break
		}

		if second_num < 1 || second_num > 10 {
			fmt.Println("Принимаем только числа от 1 до 10 ключительно!")
			break
		}

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
				fmt.Println("Деление на ноль не пройдет!")
				continue
			}
			answer = first_num / second_num
		default:
			fmt.Println("Оператор не понятный:", math_oper)
			os.Exit(1)
		}

		fmt.Printf("Ответ: %.d\n", answer)
	}
}
