package tickets

/*************************IMPORT*************************/
import (
	"errors"
	"strconv"
	"strings"
)

/*************************ESTRUCTURAS*************************/
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

/*************************METODOS*************************/
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
