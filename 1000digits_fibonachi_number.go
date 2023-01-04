//Функция позволяет определить номер числа фибоначчи, который содержит определеенное количество знаков
//в данном случае 1000 знаковое число фибоначчи имеет номер 1476

package main

import ("fmt"
	"math")


func summ_square_difference(num int) int {
	x:=0
	number1:=0.0
	number2:=1.0
	for{
	number1, number2 = number2, number1+number2
	x+=1
	if number2>=math.Pow(10.0, float64(num)){return x}
	}
}




func main (){

	fmt.Println(summ_square_difference(1000))
}
