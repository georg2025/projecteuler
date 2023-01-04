//Функция предназначена для поиска суммы цифр в числе, которое получется путем возведения двойки в определенную степень.
//например для 2 в 15 степени (32768) это будет 3+2+7+6+8=26. Для 2 в 1000 степени у меня получилось 72

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func power_digits_summ(num int) int {
	answer := 0
	n := float64(num)
	x := math.Pow(2.0, n)
	y := fmt.Sprint(x)
	r := strings.Split(y, "")
	for _, i := range r {
		s, _ := strconv.Atoi(i)
		answer += s
	}
	return answer
}

func main() {

	fmt.Println(power_digits_summ(1000))
}
