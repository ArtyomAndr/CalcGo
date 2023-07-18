package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Введите выражение:")
	reader := bufio.NewReader(os.Stdin)
	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSpace(expression)

	calculator := NewCalculator()
	result, err := calculator.EvaluateExpression(expression)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}
}

type Calculator struct {
	romanNumerals map[string]int
}

func NewCalculator() *Calculator {
	romanNumerals := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	return &Calculator{romanNumerals: romanNumerals}
}

func (c *Calculator) EvaluateExpression(expression string) (string, error) {
	// Разбиваем строку на операнды и оператор
	elements := strings.Split(expression, " ")
	if len(elements) != 3 {
		return "", errors.New("Некорректный формат математической операции")
	}

	operand1 := elements[0]
	operator := elements[1]
	operand2 := elements[2]

	// Проверяем, являются ли оба операнда арабскими или римскими цифрами
	isArabic1 := isArabicNumeral(operand1)
	isArabic2 := isArabicNumeral(operand2)
	isRoman1 := isRomanNumeral(operand1)
	isRoman2 := isRomanNumeral(operand2)

	if (isArabic1 && isRoman2) || (isRoman1 && isArabic2) {
		return "", errors.New("Используются одновременно разные системы счисления")
	}

	var result int
	var err error

	// Выполняем операцию между операндами
	if isArabic1 && isArabic2 {
		num1, _ := strconv.Atoi(operand1)
		num2, _ := strconv.Atoi(operand2)

		if !isValidArabicNumeral(num1) || !isValidArabicNumeral(num2) {
			return "", errors.New("Некорректные числа")
		}

		result, err = c.evaluateArabicExpression(num1, operator, num2)
	} else if isRoman1 && isRoman2 {
		num1 := c.convertRomanToArabic(operand1)
		num2 := c.convertRomanToArabic(operand2)

		if !isValidArabicNumeral(num1) || !isValidArabicNumeral(num2) {
			return "", errors.New("Некорректные числа")
		}

		result, err = c.evaluateArabicExpression(num1, operator, num2)
	} else {
		return "", errors.New("Некорректные числа")
	}

	if err != nil {
		return "", err
	}

	// Проверяем, в какой системе счисления были введены операнды
	if isArabic1 && isArabic2 {
		return strconv.Itoa(result), nil
	} else {
		romanResult := c.convertArabicToRoman(result)
		if romanResult == "" {
			return "", errors.New("Некорректный результат в римской системе счисления")
		}
		return romanResult, nil
	}
}

func (c *Calculator) evaluateArabicExpression(num1 int, operator string, num2 int) (int, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, errors.New("Деление на ноль")
		}
		return num1 / num2, nil
	default:
		return 0, errors.New("Некорректный оператор")
	}
}

func (c *Calculator) convertRomanToArabic(romanNumeral string) int {
	result := 0
	for i := 0; i < len(romanNumeral); i++ {
		if i > 0 && c.romanNumerals[string(romanNumeral[i])] > c.romanNumerals[string(romanNumeral[i-1])] {
			result += c.romanNumerals[string(romanNumeral[i])] - 2*c.romanNumerals[string(romanNumeral[i-1])]
		} else {
			result += c.romanNumerals[string(romanNumeral[i])]
		}
	}
	return result
}

func (c *Calculator) convertArabicToRoman(arabicNumeral int) string {
	if arabicNumeral <= 0 {
		return ""
	}

	romanNumeral := ""
	for arabicNumeral > 0 {
		if arabicNumeral >= 1000 {
			romanNumeral += "M"
			arabicNumeral -= 1000
		} else if arabicNumeral >= 900 {
			romanNumeral += "CM"
			arabicNumeral -= 900
		} else if arabicNumeral >= 500 {
			romanNumeral += "D"
			arabicNumeral -= 500
		} else if arabicNumeral >= 400 {
			romanNumeral += "CD"
			arabicNumeral -= 400
		} else if arabicNumeral >= 100 {
			romanNumeral += "C"
			arabicNumeral -= 100
		} else if arabicNumeral >= 90 {
			romanNumeral += "XC"
			arabicNumeral -= 90
		} else if arabicNumeral >= 50 {
			romanNumeral += "L"
			arabicNumeral -= 50
		} else if arabicNumeral >= 40 {
			romanNumeral += "XL"
			arabicNumeral -= 40
		} else if arabicNumeral >= 10 {
			romanNumeral += "X"
			arabicNumeral -= 10
		} else if arabicNumeral >= 9 {
			romanNumeral += "IX"
			arabicNumeral -= 9
		} else if arabicNumeral >= 5 {
			romanNumeral += "V"
			arabicNumeral -= 5
		} else if arabicNumeral >= 4 {
			romanNumeral += "IV"
			arabicNumeral -= 4
		} else {
			romanNumeral += "I"
			arabicNumeral -= 1
		}
	}
	return romanNumeral
}

func isArabicNumeral(num string) bool {
	n, err := strconv.Atoi(num)
	if err != nil {
		return false
	}
	return isValidArabicNumeral(n)
}

func isValidArabicNumeral(num int) bool {
	return num >= 1 && num <= 10
}

func isRomanNumeral(num string) bool {
	_, ok := map[string]bool{
		"I":    true,
		"II":   true,
		"III":  true,
		"IV":   true,
		"V":    true,
		"VI":   true,
		"VII":  true,
		"VIII": true,
		"IX":   true,
		"X":    true,
	}[num]
	return ok
}
