//Функция позволяет высчитать сумму чисел, которые меньше определенного значения, и, которые без остатка делятся на 3 или 5. В случае, если число делится и на 3 и на 5 - оно считается 1 раз
//Сумма таких чисел для значения до 1000 (не включительно) составляет 233 168.
package main

import "fmt"

func mult3_5(num int) int {
	answer := 0
	numbersset := make(map[int]int)
	for i := 0; i < num; i++ {
		if i%3 == 0 {
			numbersset[i] = 1
		}
		if i%5 == 0 {
			numbersset[i] = 1
		}
	}
	for f, _ := range numbersset {
		answer += f
	}

	return answer

}

func main() {

	fmt.Println(mult3_5(1000))
}
