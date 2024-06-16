package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Проект для Каты.")
	fmt.Println("Введите выражение. Можно использовать как арабские цифры, так и римские.")

	forenter := bufio.NewReader(os.Stdin)
	enter, panic := forenter.ReadString('\n')
	if panic != nil {
		fmt.Println("Ошибка чтения: ", panic)
		return
	}

	enter = strings.TrimSpace(enter)
	partsofoper := strings.Split(enter, " ")
	if len(partsofoper) != 3 {
		fmt.Println("Вы неверно ввели предложение, правильным будет ввод вида 2 + 3!")
		return
	}

	first_num, panic := strconv.ParseFloat(partsofoper[0], 64)
	if panic != nil {
		fmt.Println("Преобразование первого числа не удалось!", panic)
		return
	}

	second_num, panic := strconv.ParseFloat(partsofoper[0], 64)
	if panic != nil {
		fmt.Println("Преобразование второго числа не удалось!", panic)
		return
	}

	oper := partsofoper[1]
	var res float64

	switch oper {
	case "+":
		res = first_num + second_num
	case "-":
		res = first_num - second_num
	case "*":
		res = first_num * second_num
	case "/":
		if second_num == 0 {
			fmt.Println("Деление на ноль нельзя провернуть!")
			return
		}
		res = first_num / second_num
	default:
		fmt.Println("Вы ввели что то не понятное!")
		return

	}

	fmt.Println("Ответ: %.2f\n", res)

}
