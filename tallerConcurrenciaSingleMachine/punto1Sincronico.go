package main

import "fmt"

func proceso1(x int, y int, p12_x chan<- int, p12_y chan<- int, p21_x <-chan int, p21_y <-chan int) {

	for {

		fmt.Println("Proceso 1 Trabajando!")

		if y != 0 {

			//Envío sincrónico
			p12_x <- x
			p12_y <- y

			fmt.Println("Envío realizado 1 !")

			//Recepción sincrónica
			y = <-p21_x
			x = <-p21_y

			fmt.Println("Recepción realizada 1 !")

		} else {

			fmt.Println("El GCD es -> ", x)
			close(p12_x)
			close(p12_y)
			break

		}

	}

}

func proceso2(p21_x chan<- int, p21_y chan<- int, p12_x <-chan int, p12_y <-chan int) {

	var u, v, a int

	for {

		fmt.Println("****** 2 Trabajando!")

		//Recepción sincrónica
		u = <-p12_x
		v = <-p12_y

		fmt.Println("Recibido en 2: u=", u, "v=", v)

		/*
			//Revisar criterio de parada de la goroutine
			if u > 0 && v == 0 {

				//Cierro canales de envío
				close(p2_ex)
				close(p2_ey)

				//Terminar bucle de la goroutine
				break
			}
		*/

		a = u

		for {

			a = a - v
			if a < 0 {
				break
			} else {
				u = a
			}

			fmt.Println("P2 Enviando!!")

			//Envío sincrónico
			p21_y <- u
			p21_x <- v

			fmt.Println("P2 Finaliza envío!!")

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

	var parar string
	fmt.Scan(&parar)

}
