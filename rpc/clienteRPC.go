package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct {
	Nombre string
	Cuerpo string
}

func main() {

	var respuesta Item
	var baseDatos []Item

	conexion, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Error de conexión: ", err)
	}

	//fmt.Println("Base de Datos Inicial", baseDatos)
	a := Item{"primero", "item de prueba 1"}
	b := Item{"segundo", "item de prueba 2"}
	c := Item{"tercero", "item de prueba 3"}

	//Create
	conexion.Call("API.InsertarItem", a, &respuesta)
	conexion.Call("API.InsertarItem", b, &respuesta)
	conexion.Call("API.InsertarItem", c, &respuesta)
	conexion.Call("API.ObtenerBaseDatos", "", &baseDatos)
	fmt.Println("Segunda versión de la BD: ", baseDatos)

	//Delete
	conexion.Call("API.EliminarItem", b, &respuesta)
	conexion.Call("API.ObtenerBaseDatos", "", &baseDatos)
	fmt.Println("Segunda versión de la BD: ", baseDatos)

	//Update
	conexion.Call("API.ActualizarItem", Item{"primero", "Nueva versión del primero"}, &respuesta)
	conexion.Call("API.ObtenerBaseDatos", "", &baseDatos)
	fmt.Println("Tercera versión de la BD: ", baseDatos)

	//Read -> Consultas
	conexion.Call("API.ConsultarPorNombre", "tercero", &respuesta)
	fmt.Println("Tercer ítem: ", respuesta)

}
