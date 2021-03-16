//Ingresar el nombre de 5 productos en un array y sus respectivos precios en otro array paralelo de tipo float32.
//Mostrar cuantos productos tienen un precio mayor al primer producto ingresado (se debe contar)
package main

import "fmt"

func main() {
    var precio [5]float32
    var productos [5]string
	var suma float32

    for i := 0 ; i < len(productos); i++ {
        fmt.Println("\nIngrese el nombre del producto:")
        fmt.Scan(&productos[i])
        fmt.Println("Ingrese el precio: ")
        fmt.Scan(&precio[i])
		suma = suma + precio[i]
    }

	promedio := suma / 5
    fmt.Println("\nPromedio de los precios ",promedio)
    fmt.Println("\nProductos con precio superior al promedio ")
    for i := 0; i < len(precio); i++ {
        if precio[i] > promedio {
            fmt.Print(productos[i]," ")
        }
    }

    fmt.Println("\n\nProductos con precio superior al primer producto ")
    p1 := precio[0]
    cant := 0;

    for i := 0; i < len(precio); i++ {
        if precio[i] > p1 {
            fmt.Print(productos[i]," ")
            cant++
        }
    }

    fmt.Println("\n\nCantidad de productos con precio superior al primero: \n",cant)
}