package main

import "fmt"

func positivosNegativos() (int, int) {
    var v int
    positivos := 0
    negativos := 0
    for {        
        fmt.Print("Ingrese un numero, escriba 0 para finalizar: ")
        fmt.Scan(&v)
        if v == 0 {
            break
        } else {
            if v > 0 {
                positivos++
            } else {
                negativos++
            }
        }
    }
    return positivos, negativos
}


func main() {
    positivos, negativos := positivosNegativos()
    fmt.Println("La cantidad de los numeros positivos ingresados son:", positivos)
    fmt.Println("La cantidad de los numeros negativos ingresados son:", negativos)
}
