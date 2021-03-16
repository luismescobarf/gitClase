/* Definir un arreglo de 5 componentes de tipo float32 que representen las alturas de
5 personas. Obtener el promedio de las mismas. Contar cuántas personas son más
altas que el promedio y cuántas más bajas */
package main
import "fmt"

func main(){
	var alt [5] float32
	var prom, min, max float32
	for  i := 0; i < len(alt); i++{
		fmt.Scan(&alt[i])
		prom+=alt[i]
	}
	prom=prom/5
	for _, num := range alt {
		if(num>prom){
			max++
		} else if (num<prom){
			min++
		}
	}
	fmt.Println("Promedio:",prom,"cantidad de personas mas altas que el promedio:", max, "cantidad de personas mas bajas que el promedio:", min)

}