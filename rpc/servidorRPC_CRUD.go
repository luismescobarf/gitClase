package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	Nombre string
	Cuerpo string
}

var baseDatos []Item

type API int

//Cumple con los lineamientos de Método RPC
func (a *API) ObtenerBaseDatos(cadenaVacia string, respuesta *[]Item) error {
	*respuesta = baseDatos
	return nil
}

func (a *API) ConsultarPorNombre(nombre string, respuesta *Item) error {
	var itemObtenido Item
	for _, registro := range baseDatos {
		if registro.Nombre == nombre {
			itemObtenido = registro
		}
	}
	*respuesta = itemObtenido
	return nil
}

func (a *API) InsertarItem(item Item, respuesta *Item) error {
	baseDatos = append(baseDatos, item)
	*respuesta = item
	return nil
}

func (a *API) ActualizarItem(nuevaVersion Item, respuesta *Item) error {
	var itemCambiado Item
	for i, registro := range baseDatos {
		if registro.Nombre == nuevaVersion.Nombre {
			baseDatos[i] = nuevaVersion
			itemCambiado = nuevaVersion
		}
	}
	*respuesta = itemCambiado
	return nil
}

func (a *API) EliminarItem(item Item, respuesta *Item) error {
	var eliminado Item

	for i, registro := range baseDatos {
		if registro.Nombre == item.Nombre && registro.Cuerpo == item.Cuerpo {
			baseDatos = append(baseDatos[:i], baseDatos[i+1:]...)
			eliminado = item
			break
		}
	}
	*respuesta = eliminado
	return nil

}

func main() {

	//Registro de la API (métodos internos)
	api := new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("Error registrando la API al RPC", err)
	}

	//Iniciamos manejo de solicitudes
	rpc.HandleHTTP()

	//Inicializar la escucha de conexiones
	l, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("Error inicializando el listener", err)
	}
	log.Printf("Sirviendo RPC en el puerto %d", 4040)
	http.Serve(l, nil)

	//Error atendiendo solicitudes
	if err != nil {
		log.Fatal("Error atendiendo solicitudes!", err)
	}

	/*
		fmt.Println("Base de Datos Inicial", baseDatos)
		a := Item{"primero", "item de prueba 1"}
		b := Item{"segundo", "item de prueba 2"}
		c := Item{"tercero", "item de prueba 3"}
		//Create
		InsertarItem(a)
		InsertarItem(b)
		InsertarItem(c)
		fmt.Println("Segunda versión de la BD: ", baseDatos)
		//Delete
		EliminarItem(b)
		fmt.Println("Tercera versión de la BD: ", baseDatos)
		//Update
		ActualizarItem("tercero", Item{"cuarto", "item nueva version"})
		fmt.Println("Cuarta versión de la BD: ", baseDatos)
		//Read -> Consultas
		x := ConsultarPorNombre("cuarto")
		y := ConsultarPorNombre("primero")
		fmt.Println("Extraídos de la base de datos:")
		fmt.Println(x, y)
	*/

}
