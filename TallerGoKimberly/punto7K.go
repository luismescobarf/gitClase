//Definir un arreglo de 5 componentes de tipo float32 que representen las alturas de 5 personas. 
//Obtener el promedio de las mismas. Contar cuántas personas son más altas que el promedio y cuántas más bajas
package main

import "fmt"

func main() {
    var altura [5]float32
    var suma float32
    for i := 0; i < 5; i++ {
        fmt.Println("Ingrese la altura de la persona: ")
        fmt.Scan(&altura[i])
        suma = suma + altura[i]
    }
    promedio := suma / 5
    fmt.Println("\nPromedio de las alturas: ", promedio)
    altas := 0
    bajas := 0
    
	for i := 0; i < 5; i++ {
        if altura[i] > promedio {
            altas++
        } else {
            if altura[i] < promedio {
                bajas++
            }
        }        
    }

    fmt.Println("\nTotal de personas con altura mayor al promedio:", altas)
    fmt.Println("\nTotal de personas con altura menor al promedio:", bajas)    
}
