//Задача была найти максимальное число, которое является палиндромом (читается одинаково и справа налево и слева направо) и, при этом, ялвяется
//множеством двух трехзначных чисел. Для этого сделал небольшой скрипт для поиска наибольшего числа-палиндрома в заданных пределах
//конечно, его можно оптимизировать. Но, честно говоря, лень этим заниматься. Для конкретной задачи скорости работы хватило за глаза.


package main

import (
	"fmt"
	"strings"
)

func isPalindrom(num int) bool {
	y := fmt.Sprint(num)
	r := strings.Split(y, "")
	for i := 0; i < (len(r) / 2); i++ {
		if r[i] != r[len(r)-1-i] {
			return false
		}
	}
	return true

}

func Largest_palindrom_product(num1, num2 int) int {
	answer := 0
	result := 0
	for i := num1; i >= num2; i-- {
		for j := num1; j >= num2; j-- {
			result = i * j
			if isPalindrom(result) {
				if result > answer {
					answer = result
				}
			}
		}
	}

	return answer

}

func main() {

	fmt.Println(Largest_palindrom_product(999, 100))
}
