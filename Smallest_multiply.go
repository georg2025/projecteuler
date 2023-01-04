//Функция позволяет найти минимальное натуральное число, которое делится без остатка на все натуральные числа
//с 1 до определенного предела. В функции нет защиты от попадаения туда числа 1 или меньше. Если надо - легко могу дописать


package main

import "fmt"

func smallest_multiply(num int) int {
  y:=0
  n:=num
  if num%2!=0{n+=1}
  for{
    for i:=1; i<=num; i++{
      if n%i==0{y+=1}
    }
    if y==num{return n}
    n+=2
    y=0
  }
}

func main (){

	fmt.Println(smallest_multiply(20))
}
