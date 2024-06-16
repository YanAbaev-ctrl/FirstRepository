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

	fmt.Println("Простой консольный калькулятор")
	fmt.Println("Введите выражение (например, 2 + 3):")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		fmt.Println("Неверный формат ввода. Пример: 2 + 3")
		return
	}

	first_num, panic := strconv.ParseFloat(parts[0], 64)
	if panic != nil {
		fmt.Println("Ошибка преобразования первого числа:", panic)
		return
	}

	second_num, panic := strconv.ParseFloat(parts[2], 64)
	if panic != nil {
		fmt.Println("Ошибка преобразования второго числа:", panic)
		return
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
			fmt.Println("Деление на ноль невозможно")
			return
		}
		answer = first_num / second_num
	default:
		fmt.Println("Неподдерживаемый оператор:", math_oper)
		return
	}

	fmt.Printf("Результат: %.2f\n", answer)
}
