package tickets

type Ticket struct {
	Id            string
	Nombre        string
	Email         string
	PaisDeDestino string
	HoraVuelo     string
	Precio        string
}

// ejemplo 1
func (ticket Ticket) GetTotalTickets(destination string) (int, error) {
	var cantPersonas int
	if destination != "" && destination == ticket.{

	}
	return 0, nil
}

// ejemplo 2
func GetMornings(time string) (int, error) {
	return 0, nil

}

// ejemplo 3
func AverageDestination(destination string, total int) (int, error) {
	return 0, nil

}

/*
package tickets

type Ticket struct {
}

// ejemplo 1
func GetTotalTickets(destination string) (int, error) {}

// ejemplo 2
func GetMornings(time string) (int, error) {}

// ejemplo 3
func AverageDestination(destination string, total int) (int, error) {}
*/