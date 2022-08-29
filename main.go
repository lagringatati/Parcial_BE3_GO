package main

/*************************IMPORT*************************/
import (
	"fmt"
	"log"
	"os"
	"strings"

	packageImport "github.com/lagringatati/Parcial_BE3_GO/internal/tickets"
)

/*************************FUNCIONES*************************/
// funcion para levantar los datos de un archivo
func leerArchivo(path string) (packageImport.Tickets, error) {
	tickets := []packageImport.Ticket{}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	res, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Error al leer el archivo")
	}

	// separo los tickets por salto de linea (horizontal)
	rowsData := strings.Split(string(res), "\n")

	// por cada ticket seteo sus campos separados por coma (vertical)
	for _, rowData := range rowsData {
		if rowData != "" { //valido que no tenga líneas en blanco (sin datos)
			objectData := strings.Split(rowData, ",")
			ticketCSV := packageImport.Ticket{
				Id:            objectData[0],
				Nombre:        objectData[1],
				Email:         objectData[2],
				PaisDeDestino: objectData[3],
				HoraDelVuelo:  objectData[4],
				Precio:        objectData[5],
			}
			tickets = append(tickets, ticketCSV)
		}
	}

	return packageImport.Tickets{Tickets: tickets}, nil
}

func main() {

	data, err := leerArchivo("./tickets.csv")
	if err != nil {
		panic("Error al leer el archivo")
	}

	// Requerimiento 1 --> método "GetTotalTickets"
	/**************************************************/
	/*Posibles escenarios: 						   	  */
	/*	lugarDestino = ""			(vacio)		      */
	/*	lugarDestino = "Argentina"	(existente)	      */
	/*	lugarDestino = "dsada"		(inexistente)     */
	/**************************************************/
	/*
		var lugarDestino = ""
		totalPersonasPorDestino, errorTotalTicketsPorDestino := data.GetTotalTickets(lugarDestino)

		if errorTotalTicketsPorDestino != nil {
			fmt.Print(errorTotalTicketsPorDestino)
		} else {
			if totalPersonasPorDestino == 0 {
				fmt.Print("No hay personas que viajen a", lugarDestino)
			} else {
				fmt.Print("Las personas que viajan a", lugarDestino, "son", totalPersonasPorDestino)
			}
		}
	*/

	// Requerimiento 2 --> método "GetCountByPeriod"
	/**************************************************/
	/*Posibles escenarios: 						   	  */
	/*	periodoViaje = "dasdas"		(inexistente)     */
	/*	periodoViaje = "madrugada"	(existente->0-6)  */
	/*	periodoViaje = "mañana"		(existente->7-12) */
	/*	periodoViaje = "tarde"		(existente->13-19)*/
	/*	periodoViaje = "noche"		(existente->20-23)*/
	/**************************************************/
	/*
		var periodoViaje = "dasdas"
		totalPersonasPorPeriodo, errorPorPeriodo := data.GetCountByPeriod(periodoViaje)

		if errorPorPeriodo != nil {
			fmt.Print(errorPorPeriodo)
		} else {
			if totalPersonasPorPeriodo == 0 {
				fmt.Print("No hay personas que viajen a la", periodoViaje)
			} else {
				fmt.Print("El total de las personas que viajan a la", periodoViaje, "son", totalPersonasPorPeriodo)
			}
		}
	*/

	// Requerimiento 3 --> método "AverageDestination"
	/**************************************************/
	/*Posibles escenarios: 						   	  */
	/*	lugarDestino = ""			(vacio)		      */
	/*	lugarDestino = "Argentina"	(existente)	      */
	/*	lugarDestino = "dsada"		(inexistente)     */
	/**************************************************/
	/*
		var lugarDestino = ""
		porcentajeDePersonasPorDestino, errorPorcentajeDePersonasPorDestino := data.AverageDestination(lugarDestino)

		if errorPorcentajeDePersonasPorDestino != nil {
			fmt.Print(errorPorcentajeDePersonasPorDestino)
		} else {
			if porcentajeDePersonasPorDestino == 0 {
				fmt.Print("No hay personas que viajen a", lugarDestino)
			} else {
				fmt.Print("Las personas que viajan a", lugarDestino, "equivalen al", porcentajeDePersonasPorDestino, "% del total")
			}
		}
	*/

	// Ejecución general
	/**************************************************/
	/*Posibles escenarios: 						   	  */
	/*	lugarDestino = ""			(vacio)		      */
	/*	lugarDestino = "Argentina"	(existente)	      */
	/*	lugarDestino = "dsada"		(inexistente)     */
	/*	periodoViaje = "dasdas"		(inexistente)     */
	/*	periodoViaje = "madrugada"	(existente->0-6)  */
	/*	periodoViaje = "mañana"		(existente->7-12) */
	/*	periodoViaje = "tarde"		(existente->13-19)*/
	/*	periodoViaje = "noche"		(existente->20-23)*/
	/*	sin archivo csv o archivo vacio				  */
	/*	todas sus combinaciones						  */
	/**************************************************/

	fmt.Print("Ingrese un destino: ")
	var lugarDestino string
	fmt.Scanln(&lugarDestino)

	fmt.Print("Ingrese un período de viaje: ")
	var periodoViaje string
	fmt.Scanln(&periodoViaje)

	totalPersonasPorDestino, errorTotalTicketsPorDestino := data.GetTotalTickets(lugarDestino)
	totalPersonasPorPeriodo, errorPorPeriodo := data.GetCountByPeriod(periodoViaje)
	porcentajeDePersonasPorDestino, errorPorcentajeDePersonasPorDestino := data.AverageDestination(lugarDestino)

	fmt.Println("Ejercicio 1:")
	if errorTotalTicketsPorDestino != nil {
		fmt.Println(errorTotalTicketsPorDestino)
	} else {
		if totalPersonasPorDestino == 0 {
			fmt.Println("No hay personas que viajen a", lugarDestino)
		} else {
			fmt.Println("Las personas que viajan a", lugarDestino, "son", totalPersonasPorDestino)
		}
	}

	fmt.Println("Ejercicio 2:")
	if errorPorPeriodo != nil {
		fmt.Println(errorPorPeriodo)
	} else {
		if totalPersonasPorPeriodo == 0 {
			fmt.Println("No hay personas que viajen a la", periodoViaje)
		} else {
			fmt.Println("El total de las personas que viajan a la", periodoViaje, "son", totalPersonasPorPeriodo)
		}
	}

	fmt.Println("Ejercicio 3:")
	if errorPorcentajeDePersonasPorDestino != nil {
		fmt.Println(errorPorcentajeDePersonasPorDestino)
	} else {
		if porcentajeDePersonasPorDestino == 0 {
			fmt.Println("No hay personas que viajen a", lugarDestino)
		} else {
			fmt.Println("Las personas que viajan a", lugarDestino, "equivalen al", porcentajeDePersonasPorDestino, "% del total")
		}
	}
}
