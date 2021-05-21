package main

import "fmt"

type Item struct {
	Nombre string
	Cuerpo string
}

var baseDatos []Item

func ConsultarPorNombre(nombre string) Item {
	var itemObtenido Item
	for _, registro := range baseDatos {
		if registro.Nombre == nombre {
			itemObtenido = registro
		}
	}
	return itemObtenido
}

func InsertarItem(item Item) Item {
	baseDatos = append(baseDatos, item)
	return item
}

func ActualizarItem(nombre string, nuevaVersion Item) Item {
	var itemCambiado Item
	for i, registro := range baseDatos {
		if registro.Nombre == nombre {
			baseDatos[i] = nuevaVersion
			itemCambiado = nuevaVersion
		}
	}
	return itemCambiado
}

func EliminarItem(item Item) Item {
	var eliminado Item

	for i, registro := range baseDatos {
		if registro.Nombre == item.Nombre && registro.Cuerpo == item.Cuerpo {
			baseDatos = append(baseDatos[:i], baseDatos[i+1:]...)
			eliminado = item
			break
		}
	}
	return eliminado

}

func main() {

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

}
