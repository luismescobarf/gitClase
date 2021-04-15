// 1. Implementar la metodología presentada en clase para obtener el máximo común
// divisor, del libro (Czaja, 2018) en versión secuencial. Luego implementarlo de
// manera distribuida con comunicación sincrónica
package main 
import "fmt"
func gcd (x int, y int )int{
	for{
		if x >0 && y ==0{
		
			return x
		}else {
			r:= x%y
			x=y
			y=r
			
		}

	}
	
}
func main (){

var x,y int 
x=9
y=6

	fmt.Println("El gcd de ",x,"y ",y,"es: ",gcd(x,y))
}











