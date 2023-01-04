//Задача была посчитать все числа, кроме 1 и 2, которые равны сумме факториалов составляющих их цифр. Например число 145 = 1!+4!+5!
//для этого я написал пару небольших функций. Также, перемножил 9! на 6 и получил чуть больше 2 миллионов. В итоге пришел к выводу, что
//такие числа не могут быть выше 2 500 000. Поставил цикл до этого числа и кроме 145 нашел всего одно: 40585.


package main

import (
	"fmt"
	"strconv"
	"strings"
)

func factorial(num float64) float64 {
	answer := 1.0
	for i := 1; i <= int(num); i++ {
		answer *= float64(i)
	}
	return answer

}

func digits_factorial(num int) int {
	answer := 0
	answer2 := 0
	for i := 3; i <= num; i++ {
		y := fmt.Sprint(i)
		r := strings.Split(y, "")
		answer2 = 0
		for _, n := range r {
			z, _ := strconv.ParseFloat(n, 64)
			answer2 += int(factorial(z))
		}
		if answer2 == i {
			answer += i
			fmt.Println(i)
		}

	}
	return answer

}

func main() {

	fmt.Println(digits_factorial(2500000))
}
