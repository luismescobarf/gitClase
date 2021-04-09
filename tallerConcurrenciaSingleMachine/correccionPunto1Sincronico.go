package main

import "fmt"

type valores struct {
	x int
	y int
}

func proceso1(tupla valores, p12 chan valores, p21 chan valores) {

	for {

		if tupla.y != 0 {

			//Envío sincrónico
			p12_x <- x
			p12_y <- y

			////Proceso 2
			go proceso2(p12, p21)

			//Recepción sincrónica
			y = <-p21_x
			x = <-p21_y

		} else {

			fmt.Println("El GCD es -> ", x)
			break

		}

	}

}

func proceso2(p12 chan valores, p21 chan valores) {

	var u, v, a int

	tupla := <-p12
	u = tupla.x
	v = tupla.y

	a = u

	for {

		a = a - v
		if a < 0 {
			break
		} else {
			u = a
		}

	}

	//Envío sincrónico
	tupla.y = u
	tupla.x = v
	p21 <- tupla

}

func main() {

	var x, y int
	x = 9
	y = 6

	tupla := make(valores)
	tupla.x = x
	tupla.y = y

	//Diseño distribuído y sincrónico
	p12 := make(chan valores)
	p21 := make(chan valores)

	//Lanzamiento del proceso 1
	go proceso1(tupla, p12, p21)

	var parar string
	fmt.Scan(&parar)

}
