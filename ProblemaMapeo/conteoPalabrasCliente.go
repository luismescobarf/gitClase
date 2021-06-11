package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
)

type Mensaje struct {
	Mensaje string
}

//El main realizará en la versión secuencial el Input, el Splitting y el Final result
func main() {

	//inicio := time.Now() //Inicio de la toma de tiempo

	//Splitting implícito por líneas
	var ruta string

	//Entradas de prueba
	//ruta = "texto.txt"
	//ruta = "foo.txt"
	//Entrada de libros completos
	ruta = "./libros/DonQuijote.txt"
	//ruta = "./libros/Iliad.txt"
	//ruta = "./libros/AliceWonderland.txt"

	//fmt.Println(envio)
	//Armar mensaje de prueba
	mensaje := Mensaje{Mensaje: ruta}
	bufferContenedor := new(bytes.Buffer)

	c, err := net.Dial("tcp", "localhost:12345")
	if err != nil {

		fmt.Println(err)
		return

		//Alternativas del manejo
		//Panic
		//Log

	}
	defer c.Close()

	//Crear objeto codificador
	gobobj := gob.NewEncoder(bufferContenedor)

	//Codificar el buffer y encapsularlo (marshal) en un objeto gob
	gobobj.Encode(mensaje)

	//Hacer el envío
	c.Write(bufferContenedor.Bytes())

}
