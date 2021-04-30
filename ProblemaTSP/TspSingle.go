package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"time"
)

//Creamos las variables de instancia
var inicio time.Time
var fin time.Time

//Funcion para crear la matriz de adyacencia
func crearMatrizAdyacencia() map[string]int {
	matrizAdyacencia := map[string]int{
		"0-0": 0, "0-1": 21, "0-2": 57, "0-3": 58, "0-4": 195, "0-5": 245, "0-6": 241,
		"1-0": 127, "1-1": 0, "1-2": 231, "1-3": 113, "1-4": 254, "1-5": 179, "1-6": 41,
		"2-0": 153, "2-1": 252, "2-2": 0, "2-3": 56, "2-4": 126, "2-5": 160, "2-6": 269,
		"3-0": 196, "3-1": 128, "3-2": 80, "3-3": 0, "3-4": 136, "3-5": 37, "3-6": 180,
		"4-0": 30, "4-1": 40, "4-2": 256, "4-3": 121, "4-4": 0, "4-5": 194, "4-6": 109,
		"5-0": 33, "5-1": 144, "5-2": 179, "5-3": 114, "5-4": 237, "5-5": 0, "5-6": 119,
		"6-0": 267, "6-1": 61, "6-2": 79, "6-3": 39, "6-4": 135, "6-5": 55, "6-6": 0,
	}
	return matrizAdyacencia
}

//Generamos las rutas para el multiarranque
func generarRutas(matrizAdyacencia map[string]int) [4][8]string {
	var rutas [4][8]string
	for i := 0; i < 4; i++ {
		if i == 0 {
			for j := 0; j < int(math.Sqrt(float64(len(matrizAdyacencia))))+1; j++ {
				rutas[i][j] = strconv.Itoa(j)
			}
			rutas[i][int(math.Sqrt(float64(len(matrizAdyacencia))))] = strconv.Itoa(0)
		}
		if i == 1 {
			var ruta []string
			for j := 0; j < int(math.Sqrt(float64(len(matrizAdyacencia))))+1; j++ {
				if (j % 2) == 0 {
					ruta = append(ruta, strconv.Itoa(j))
				}
			}
			for j := 0; j < int(math.Sqrt(float64(len(matrizAdyacencia))))+1; j++ {
				if (j % 2) == 1 {
					ruta = append(ruta, strconv.Itoa(j))
				}
			}
			for j := 0; j < int(math.Sqrt(float64(len(matrizAdyacencia))))+1; j++ {
				rutas[i][j] = ruta[j]
			}
			rutas[i][int(math.Sqrt(float64(len(matrizAdyacencia))))] = strconv.Itoa(0)

		}
		if i == 2 {
			var ruta []string
			ruta = append(ruta, strconv.Itoa(0))
			for j := 1; j < int(math.Sqrt(float64(len(matrizAdyacencia)))); j++ {
				if (j % 2) == 1 {
					ruta = append(ruta, strconv.Itoa(j))
				}
			}
			for j := 1; j < int(math.Sqrt(float64(len(matrizAdyacencia)))); j++ {
				if (j % 2) == 0 {
					ruta = append(ruta, strconv.Itoa(j))
				}
			}
			ruta = append(ruta, strconv.Itoa(0))
			for j := 0; j < int(math.Sqrt(float64(len(matrizAdyacencia))))+1; j++ {
				rutas[i][j] = ruta[j]
			}
			rutas[i][int(math.Sqrt(float64(len(matrizAdyacencia))))] = strconv.Itoa(0)
		}
		if i == 3 {
			var ruta []string
			for j := 0; j < int(math.Sqrt(float64(len(matrizAdyacencia)))); j++ {
				ruta = append(ruta, strconv.Itoa(j))
			}
			ruta = append(ruta, strconv.Itoa(0))
			cambio := ruta[2:6]
			sort.Slice(cambio, func(i, j int) bool { return cambio[i] > cambio[j] })
			for j := 2; j < 6; j++ {
				ruta[j] = cambio[j-2]
			}
			for j := 0; j < int(math.Sqrt(float64(len(matrizAdyacencia))))+1; j++ {
				rutas[i][j] = ruta[j]
			}
		}
	}
	return rutas
}

//Generamos el costo de las rutas
func costoTour(matrizAdyacencia map[string]int, rutas [4][8]string) [4]int {
	var costos [4]int
	for i := 0; i < 4; i++ {
		var ruta []string
		for j := 0; j < 8; j++ {
			ruta = append(ruta, rutas[i][j])
		}
		llaves := llavesMatriz(ruta)
		costo := calcularCosto(matrizAdyacencia, llaves)
		costos[i] = costo
	}
	return costos
}

//Generamos el costo de las rutas
func tourOptimizado(matrizAdyacencia map[string]int, rutas [4][8]string) [4][8]string {
	var rutasO [4][8]string
	for i := 0; i < 4; i++ {
		var ruta []string
		for j := 0; j < 8; j++ {
			ruta = append(ruta, rutas[i][j])
		}
		fmt.Println("")
		fmt.Println("Optimizacion de ruta ", i+1)
		fmt.Println("")
		rutaO := optimizarRuta(matrizAdyacencia, ruta)
		for j := 0; j < 8; j++ {
			rutasO[i][j] = rutaO[j]
		}

	}
	return rutasO
}

//Creamos las cadenas para las llaves de la matriz
func llavesMatriz(matriz []string) []string {
	var llaves []string
	for i := 0; i < len(matriz)-1; i++ {
		llaves = append(llaves, matriz[i]+"-"+matriz[i+1])
	}
	return llaves
}

//Calcular costos de las rutas
func calcularCosto(matrizAdyacencia map[string]int, llave []string) int {
	var costo int
	for i := 0; i < len(llave); i++ {
		costo = costo + matrizAdyacencia[llave[i]]
	}
	return costo
}

//Generar una nueva vareacion
func optimizarRuta(matrizAdyacencia map[string]int, ruta []string) []string {

	//Definimos el costo de la ruta y las rutas a cambiar
	rutaActual := ruta
	mejorRuta := ruta
	mejorDistancia := calcularCosto(matrizAdyacencia, llavesMatriz(ruta))
	distanciaActual := 0
	distanciaNueva := mejorDistancia
	rutaNueva := make([]string, len(ruta))

	for l := 0; l < 2; l++ {

		var variable bool = true
		var mejora bool = false

		for variable {
			for i := 1; i < len(ruta)-2; i++ {
				for j := i + 1; j < len(ruta)-1; j++ {
					copy(rutaNueva, rutaActual)
					rutaNueva[i], rutaNueva[j] = rutaNueva[j], rutaNueva[i]
					distanciaActual = calcularCosto(matrizAdyacencia, llavesMatriz(rutaNueva))
					if distanciaActual < distanciaNueva {
						fmt.Println(rutaActual, "--->", rutaNueva)
						distanciaNueva = distanciaActual
						copy(mejorRuta, rutaNueva)
						distanciaActual = 0
						mejora = true
					} else {
						distanciaActual = 0
					}
				}
			}
			variable = false
			if mejora {
				copy(rutaActual, mejorRuta)
				mejorDistancia = distanciaNueva
				mejora = false
			}
		}
	}
	return rutaActual
}

//Creamos el main
func main() {

	inicio = time.Now()

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//Funcion MultiArranque
	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	//Creamos la matriz de adyacencia
	matrizAdyacencia := crearMatrizAdyacencia()
	rutas := generarRutas(matrizAdyacencia)
	costos := costoTour(matrizAdyacencia, rutas)
	fmt.Println("")
	fmt.Println("Rutas generadas con costos")
	fmt.Println("")
	for i := 0; i < len(rutas); i++ {
		fmt.Print(rutas[i])
		fmt.Println(" ", costos[i])
	}
	rutasO := tourOptimizado(matrizAdyacencia, rutas)
	costosO := costoTour(matrizAdyacencia, rutasO)
	fmt.Println("")
	fmt.Println("Rutas Optimizadas generadas con costos")
	fmt.Println("")
	for i := 0; i < len(rutas); i++ {
		fmt.Print(rutasO[i])
		fmt.Println(" ", costosO[i])
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//Funcion UnicoArranque
	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	//Creamos una posible ruta
	// matrizRuta := []string{"0", "1", "2", "3", "4", "5", "6", "0"}
	// matrizAdyacencia := crearMatrizAdyacencia()
	// fmt.Println("Traveling Salesman Problem")
	// fmt.Println(matrizAdyacencia)
	// fmt.Println(matrizRuta)
	// llaves := llavesMatriz(matrizRuta)
	// fmt.Println(llaves)
	// fmt.Println("Costo total: ", calcularCosto(matrizAdyacencia, llaves))
	// var mensaje = optimizarRuta(matrizAdyacencia, matrizRuta)
	// fmt.Println("La nueva ruta generada es: ", mensaje)
	// fmt.Println("Con un costo de: ", calcularCosto(matrizAdyacencia, llavesMatriz(mensaje)))

	fmt.Println("")

	fin = time.Now()

	//Restamos la diferencia de timepo desde que inicio hasta que finaliza
	dif := fin.Sub(inicio)
	fmt.Println("Tiempo de ejecución: ")
	fmt.Println(dif)
}
