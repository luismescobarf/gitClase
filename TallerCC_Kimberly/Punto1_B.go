// 1. Implementar la metodología presentada en clase para obtener el máximo común
// divisor, del libro (Czaja, 2018) en versión secuencial. Luego implementarlo de
// manera distribuida con comunicación sincrónica.
// 1. Implementar la metodología presentada en clase para obtener el máximo común
// divisor, del libro (Czaja, 2018) en versión secuencial. Luego implementarlo de
// manera distribuida con comunicación sincrónica
package main 
import "fmt"


func proceso1 (x int, y int p1_ex chan <= int, p1_ey chan <= int,p1_rx<=chan int,p1_ry<=chan int){
	for{
		if y != 0{

			// Envio sincronico
			p1_ex<-x
			p1_ey<-y

			// Recepcion sincronica
			y = <-p1_rx
			x = <-p1_ry
		}else{
			
			fmt.Println("El GCD es -> ",x)
			close(p1_ex)
			close(p1_ey)
			break;
		}
	}
}

func proceso2 (p2_ex chan <= int, p2_ey chan <= int,p2_rx<=chan int,p2_ry<=chan int){
	
	var u, v , a int
	for{

		// Recepcion sincronica
		u = <-p2_rx
		y = <-p2_ry


		// Criterio de parada
		if u > 0&& v=={
			// Cierro canales de envio
			close(p2_ex)
			close(p2_ey)
			break;

		}

		a=u

		for{
			a = a-v	
			if a <0{
				break
			}else{
				u=a
			}			

			// Envio sincronico
			p2_ex<-u
			p2_ey<-v
		}

	}
}

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

// Diseño distribuido sincronico
p12_x := make(chan int)
p12_y := make(chan int)
p21_y := make(chan int)
p21_x := make(chan int)

go proceso1(x,y, p12_x, p12_y, p)

}











