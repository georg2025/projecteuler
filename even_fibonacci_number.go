//Зря эту задачу решил на голанг сделать. Понял уже только когда столкнулся с тем, что очень нужен был while. Ну да ладно, вроде выкрутился
//Функция дает сумму четных чисел фибоначчи в диапазоне до заданного числа. В частности, для 4 000 000 эта сумма равняется 4 613 732

package main

import "fmt"

func even_numbers_count(num int) int {
	answer := 0
	number1 := 0
	number2 := 1
	for {
		number1, number2 = number2, number1+number2
		if number2%2 == 0 {
			answer += number2
		}
		if number2 >= num {
			return answer
		}
	}

}

func main() {

	fmt.Println(even_numbers_count(4000000))
}
