package main

import "fmt"

func proceso1(x int, y int, p1_ex chan<- int, p1_ey chan<- int, p1_rx <-chan int, p1_ry <-chan int) {

	for {

		if y != 0 {

			//Envío sincrónico
			p1_ex <- x
			p1_ey <- y

			//Recepción sincrónica
			y = <-p1_rx
			x = <-p1_ry

		} else {

			fmt.Println("El GCD es -> ", x)
			close(p1_ex)
			close(p1_ey)
			break

		}

	}

}

func proceso2(p2_ex chan<- int, p2_ey chan<- int, p2_rx <-chan int, p2_ry <-chan int) {

	var u, v, a int

	for {

		//Recepción sincrónica
		u = <-p2_rx
		v = <-p2_ry

		//Revisar criterio de parada de la goroutine
		if u > 0 && v == 0 {

			//Cierro canales de envío
			close(p2_ex)
			close(p2_ey)

			//Terminar bucle de la goroutine
			break
		}

		a = u

		for {

			a = a - v
			if a < 0 {
				break
			} else {
				u = a
			}

			//Envío sincrónico
			p2_ey <- u
			p2_ex <- v

		}

	}

}

//Obtener el maximo comun divisor single threaded
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

	//Llamado single threaded
	//fmt.Println("El GCD de ", x, " y ", y, " es -> ", gcd(x, y))

	//Diseño distribuído y sincrónico
	p12_x := make(chan int)
	p12_y := make(chan int)
	p21_x := make(chan int)
	p21_y := make(chan int)

	//Lanzamiento del proceso 1
	go proceso1(x, y, p12_x, p12_y, p21_x, p21_y)

	//Lanzamiento del proceso 2
	go proceso2(p21_x, p21_y, p12_x, p12_y)

}
