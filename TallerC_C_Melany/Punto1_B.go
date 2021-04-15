package main

import (
	"fmt"
	"sync"
)

//Estructura de sincronización de goroutines
var wg sync.WaitGroup

func proceso1(x int, y int, canalGeneral chan int) {

	//Avisar terminación de proceso 1 al goroutine main que está en espera
	defer wg.Done()

	for {

		//Salida de diagnóstico
		//fmt.Println("Proceso 1 Trabajando!")

		if y != 0 {

			//Envío sincrónico
			canalGeneral <- x
			canalGeneral <- y

			fmt.Println("Envío realizado 1 !")

			//Recepción sincrónica
			y = <-canalGeneral
			x = <-canalGeneral

			fmt.Println("Recepción realizada 1 !")

		} else {

			fmt.Println("El GCD es -> ", x)
			break

		}

	}

	fmt.Println("P1 termina!")

}

func proceso2(canalGeneral chan int) {

	//Avisar terminación de proceso 2 al goroutine main que está en espera
	defer wg.Done()

	var u, v, a int

	for {

		//Recepción sincrónica
		u = <-canalGeneral
		v = <-canalGeneral

		fmt.Println("Recibido en 2: u=", u, "v=", v)

		a = u

		for {

			a = a - v
			if a < 0 {
				break
			} else {
				u = a
			}

		}

		fmt.Println("P2 Enviando!! u=", u, "v=", v)

		//Envío sincrónico
		canalGeneral <- u
		canalGeneral <- v

		if u == 0 {
			break
		}

	}

	fmt.Println("P2 termina!")

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
	x = 60
	y = 45

	//Llamado single threaded
	//fmt.Println("El GCD de ", x, " y ", y, " es -> ", gcd(x, y))

	//Diseño distribuído y sincrónico
	canalGeneral := make(chan int)

	//Sincronización del main con las dos goroutines
	wg.Add(2)

	//Lanzamiento del proceso 1
	go proceso1(x, y, canalGeneral)

	//Lanzamiento del proceso 2
	go proceso2(canalGeneral)

	//Esperar a que los dos procesos terminen
	wg.Wait()

}