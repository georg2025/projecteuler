//Тут чисто задача на комбинаторику. Только фотруму подставить. Нужно отыскать количество возможных вариантов путей, если
//число шагов в низ равняется х, а число шагов вперед - y. Банально берем факториал суммы чисел и делим на факториалы каждого из чисел.

package main

import "fmt"


func factorial(num float64) float64{
    answer:=1.0
    for i:=1; i<=int(num); i++{
      answer*=float64(i)
    }
    return answer

}

func lattice_path(num1, num2 int) float64{
  n:=float64(num1)
  n2:=float64(num2)
  return(factorial(n+n2)/factorial(n)/factorial(n2))

}




func main (){

	fmt.Println(lattice_path(20,20))
}
