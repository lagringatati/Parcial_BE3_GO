package main

/*************************IMPORT*************************/
import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*************************ESTUCTURAS*************************/
// Slice Tickets con elementos de tipo Ticket ("herencia" --> embedding structs)
type Tickets struct {
	Tickets []Ticket
}

// Estructura de un Ticket
type Ticket struct {
	Id            string
	Nombre        string
	Email         string
	PaisDeDestino string
	HoraDelVuelo  string
	Precio        string
}

/*************************FUNCIONES/METODOS*************************/
// funcion para levantar los datos de un archivo
func leerArchivo(path string) (Tickets, error) {
	tickets := []Ticket{}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	res, _ := os.ReadFile(path)

	// separo los tickets por salto de linea (horizontal)
	rowsData := strings.Split(string(res), "\n")

	// por cada ticket seteo sus campos separados por coma (vertical)
	for _, rowData := range rowsData {
		if rowData != "" { //valido que no tenga líneas en blanco (sin datos)
			objectData := strings.Split(rowData, ",")
			ticketCSV := Ticket{
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

	return Tickets{Tickets: tickets}, nil
}

/*************************REQUERIMIENTO 1*************************/
// método "GetTotalTickets"
func (tickets Tickets) GetTotalTickets(destination string) (int, error) {
	var totalPersonasQueViajan int
	totalPersonasQueViajan = 0

	if destination == "" {
		return totalPersonasQueViajan, errors.New("No se ingresó ningún lugar de destino")
	}

	for _, ticket := range tickets.Tickets { // range <= []Ticket
		if destination == ticket.PaisDeDestino {
			totalPersonasQueViajan++
		}
	}

	return totalPersonasQueViajan, nil
}

/*************************REQUERIMIENTO 2*************************/
// funcion auxiliar
func cantidadDePersonasPorRango(tickets Tickets, horaInicio, horaFin int64) (int, error) {
	var totalPersonasPorRango int
	totalPersonasPorRango = 0
	if horaInicio < 0 || horaInicio > 23 || horaFin < 0 || horaFin > 23 || horaInicio >= horaFin {
		return totalPersonasPorRango, errors.New("El rango de horas ingresado es inválido")
	}
	for _, ticket := range tickets.Tickets {
		// separo la hora exacta del vuelo en horas y minutos (mediante ":")
		horaExactaSplit := strings.Split(ticket.HoraDelVuelo, ":")
		hora, _ := strconv.ParseInt(horaExactaSplit[0], 0, 64)

		if hora >= horaInicio && hora <= horaFin { // valido que se encuentre en el rango establecido
			totalPersonasPorRango++
		}
	}

	return totalPersonasPorRango, nil
}

// método "GetCountByPeriod"
func (tickets Tickets) GetCountByPeriod(time string) (int, error) {
	switch strings.ToLower(time) {
	case "madrugada":
		return cantidadDePersonasPorRango(tickets, 0, 6)
	case "mañana":
		return cantidadDePersonasPorRango(tickets, 7, 12)
	case "tarde":
		return cantidadDePersonasPorRango(tickets, 13, 19)
	case "noche":
		return cantidadDePersonasPorRango(tickets, 20, 23)
	default:
		return 0, errors.New("Período inexistente de viaje")
	}
}

/*************************REQUERIMIENTO 3*************************/
// método "AverageDestination" (devuelve el porcentaje de personas que viajan a un destino determinado)
func (tickets Tickets) AverageDestination(destination string) (float64, error) {
	var totalPersonas int
	var totalPorcentaje float64
	totalPersonasPorDestino, errorTotalTicketsPorDestino := tickets.GetTotalTickets(destination)
	totalPersonas = len(tickets.Tickets)
	if errorTotalTicketsPorDestino != nil {
		return 0, errorTotalTicketsPorDestino
	}
	if totalPersonas == 0 { // valido que no se divida por cero
		return 0, errors.New("No hay ningún ticket vendido aún")
	}

	totalPorcentaje = (float64(totalPersonasPorDestino) / float64(totalPersonas)) * 100

	return totalPorcentaje, nil
}

func main() {

	data, err := leerArchivo("./tickets.csv")
	if err != nil {
		fmt.Println("Error en la lectura del archivo")
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
				fmt.Print("No hay personas que viajen a ", lugarDestino)
			} else {
				fmt.Print("Las personas que viajan a ", lugarDestino, " son ", totalPersonasPorDestino)
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
		var periodoViaje = ""
		totalPersonasPorPeriodo, errorPorPeriodo := data.GetCountByPeriod(periodoViaje)

		if errorPorPeriodo != nil {
			fmt.Print(errorPorPeriodo)
		} else {
			if totalPersonasPorPeriodo == 0 {
				fmt.Print("No hay personas que viajen a la ", periodoViaje)
			} else {
				fmt.Print("Las personas que viajan a la ", periodoViaje, " son ", totalPersonasPorPeriodo)
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
				fmt.Print("No hay personas que viajen a ", lugarDestino)
			} else {
				fmt.Print("Las personas que viajan a ", lugarDestino, " equivalen al ", porcentajeDePersonasPorDestino, " porciento del total")
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

	if errorTotalTicketsPorDestino != nil {
		fmt.Println(errorTotalTicketsPorDestino)
	} else {
		if totalPersonasPorDestino == 0 {
			fmt.Println("No hay personas que viajen a ", lugarDestino)
		} else {
			fmt.Println("Las personas que viajan a ", lugarDestino, " son ", totalPersonasPorDestino)
		}
	}

	if errorPorPeriodo != nil {
		fmt.Println(errorPorPeriodo)
	} else {
		if totalPersonasPorPeriodo == 0 {
			fmt.Println("No hay personas que viajen a la ", periodoViaje)
		} else {
			fmt.Println("Las personas que viajan a la ", periodoViaje, " son ", totalPersonasPorPeriodo)
		}
	}

	if errorPorcentajeDePersonasPorDestino != nil {
		fmt.Println(errorPorcentajeDePersonasPorDestino)
	} else {
		if porcentajeDePersonasPorDestino == 0 {
			fmt.Println("No hay personas que viajen a ", lugarDestino)
		} else {
			fmt.Println("Las personas que viajan a ", lugarDestino, " equivalen al ", porcentajeDePersonasPorDestino, " por ciento del total")
		}
	}

}
