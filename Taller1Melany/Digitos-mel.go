package  main 
import "fmt"

func main(){
var n1 int 
fmt.Print("Ingrese el numero positivo del 1 al 99: ")
fmt.Scan(&n1)


if n1>=100{
	
fmt.Println("Es del 1 al 99 :| ")

}else {
	if n1<100&&n1>=10{
		fmt.Println("Este numero contiene dos digitos ")
	}else{
		fmt.Println("Este numero contiene un digito ")	}
}


}


