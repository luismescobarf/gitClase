package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
)

type Mensaje struct {
	Id       string
	Mensaje  string
	Promedio float32
}

func main() {

	//Armar mensaje de prueba
	mensaje := Mensaje{Id: "AA234", Mensaje: "Hola Mundo Estructura Desde Cliente Marshalled", Promedio: 4.5}
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
