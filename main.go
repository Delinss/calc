package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Введите выражение: (например, \"hello\" + \"world\")")

	// Используем bufio.Scanner для ввода всей строки
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Ошибка:", r)
			}
		}()

		result := calculate(input)
		fmt.Println("Результат:", truncateResult(result))
	} else {
		fmt.Println("Ошибка при вводе данных.")
	}
}

func calculate(input string) string {
	input = strings.TrimSpace(input)

	// Поддерживаемые операции
	operations := []string{"+", " - ", "*", "/"}
	var operation string

	for _, op := range operations {
		if strings.Contains(input, op) {
			operation = op
			fmt.Println("op = ", op)
			break
		}
	}
	fmt.Println("operation = ", operation)

	if operation == "" {
		panic("Неподдерживаемая операция")
	}
	fmt.Println("operation = ", operation)

	parts := strings.Split(input, operation)
	if len(parts) != 2 {
		panic("Некорректный формат ввода")
	}

	left := strings.TrimSpace(parts[0])
	right := strings.TrimSpace(parts[1])

	if !strings.HasPrefix(left, "\"") || !strings.HasSuffix(left, "\"") {
		panic("Первым аргументом должна быть строка в кавычках")
	}

	left = left[1 : len(left)-1] // Убираем кавычки

	if len(left) > 10 {
		panic("Длина строки не должна превышать 10 символов")
	}
	fmt.Println("operation = ", operation)
	if operation == "+" || operation == " - " {
		if !strings.HasPrefix(right, "\"") || !strings.HasSuffix(right, "\"") {
			panic("Вторым аргументом должна быть строка в кавычках")
		}
		right = right[1 : len(right)-1]

		if len(right) > 10 {
			panic("Длина строки не должна превышать 10 символов")
		}

		if operation == "+" {
			return left + right
		}

		if operation == " - " {
			if strings.Contains(left, right) {
				return strings.Replace(left, right, "", 1)
			}
			return left
		}
	}

	// Операции умножения и деления
	num, err := strconv.Atoi(right)
	if err != nil || num < 1 || num > 10 {
		panic("Число должно быть от 1 до 10")
	}

	if operation == "*" {
		return strings.Repeat(left, num)
	}

	if operation == "/" {
		segmentLength := len(left) / num
		return left[:segmentLength]
	}

	panic("Неподдерживаемая операция")
}

func truncateResult(result string) string {
	if len(result) > 40 {
		return result[:40] + "..."
	}
	return result
}
