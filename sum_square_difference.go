//Функция выполняет следующие действия: берет сумму чисел от 1 до n и возводит эту сумму в квадрат. Потом из полученного результата
//вычитает сумму квадратов чисел от 1 до n


package main

import ("fmt"
  "math")

func summ_square_difference(num int) int {
  y:=0
  for i:=1; i<=num; i++{
    y+=int(math.Pow(float64(i),2))
  }
  x:=0
  for i:=1; i<=num; i++{
    x+=i
  }
  return(int(math.Pow(float64(x),2))-y)

}

func main (){

	fmt.Println(summ_square_difference(100))
}
