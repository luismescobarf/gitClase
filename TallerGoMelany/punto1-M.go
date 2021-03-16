package  main 
import "fmt"

func main(){
var n1, n2 string 
fmt.Print("Ingrese el primer nombre: ")
fmt.Scan(&n1)

fmt.Print("Ingrese el segundo nombre: ")
fmt.Scan(&n2)

if n1<n2{
	
fmt.Println(n1,"\n",n2)

}else {
	fmt.Println(n2,"\n",n1)
}

}

