package main

import "fmt"

//Obtener el maximo comun divisor
func gcd(x int, y int) int {

	for {
		if x > 0 && y == 0 {
			return x
		} else {
			r := x % y
			x = y
			y = r
		}
	}

}

func main() {

	var x, y int
	x = 9
	y = 6
	fmt.Println("El GCD de ", x, " y ", y, " es -> ", gcd(x, y))

}
