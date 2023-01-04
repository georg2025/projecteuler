//Функция находит наибольший простой делитель для определенного числа. В ней, правда, нет защиты от просто числа в условиях
//Если надо так-то могу дописать. Сейчас, если передать в аргумент функции простое число - будет ошибка деления на 0


package main

import "fmt"

func largest_prime_factor(num int) int {

  y:=num/2
  x:=0
  if num<=0{return 0}
  if num==1{return 1}
  for{

    if num%y==0{
      x=0
      for i:=1; i<=y; i++{
        if y%i==0{x+=1}
        if x>2{break}
      }
      if x==2{return y}

    }
    y--
  }
}

func main (){

	fmt.Println(largest_prime_factor(600851475143))
}
