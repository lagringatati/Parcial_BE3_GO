package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Tickets struct{
	Tickets []Ticket
}

type Ticket struct {
	Id            string
	Nombre        string
	Email         string
	PaisDeDestino string
	HoraVuelo     string
	Precio        string
}

// ejemplo 1
func (tickets Tickets) GetTotalTickets(destination string) (int, error) {
	var cantPersonas int
	if destination == ""{
		return 0, errors.New("")
	}

	for _, ticket := range tickets.Tickets {
		if destination == ticket.PaisDeDestino{
			cantPersonas++
		}
	}

	return cantPersonas, nil
}

func calcularCantidadDePersonas( tickets Tickets, desde, hasta int64) (int, error){
	var cantPersonas int

	for _, ticket := range tickets.Tickets {
		horaVueloSlice := strings.Split(ticket.HoraVuelo, ":")
		horaVuelo,_ := strconv.ParseInt(horaVueloSlice[0], 0, 64)
	
		if desde <= horaVuelo && hasta >= horaVuelo{
			cantPersonas ++
		}
	}

	return cantPersonas, nil
}

// ejemplo 2
func (tickets Tickets) GetMornings(time string) (int, error) {
	switch strings.ToLower(time) {
	case "madrugada":
		return calcularCantidadDePersonas(tickets, 0, 6)
	case "mañana":
		return calcularCantidadDePersonas(tickets, 7, 12)
	case "tarde":
		return calcularCantidadDePersonas(tickets, 13, 19)
	case "noche":
		return calcularCantidadDePersonas(tickets, 20, 23)
	default:
		return 0, errors.New("no existe el periodo de tiempo indicado")
	}
}

// ejemplo 3
func (tickets Tickets) AverageDestination(destination string) (float64, error) {
	var totalPersonas float64
	var totalPersonasDestino float64

	if destination == ""{
		return 0, errors.New("kaks")
	}

	for _, ticket := range tickets.Tickets {
		totalPersonas++
		if destination == ticket.PaisDeDestino{
			totalPersonasDestino++
		}
	}

	if totalPersonasDestino == 0 {
		return 0, errors.New("no existen tickets para este país")
	}

	fmt.Println(totalPersonas)
	fmt.Println(totalPersonasDestino)

	porcentaje := math.Round(totalPersonasDestino/totalPersonas*100)
	return porcentaje, nil
}


func readFile(path string) (Tickets, error) {
	tickets := []Ticket{}
	
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	res, _ := os.ReadFile(path)
	
	rowsData := strings.Split(string(res), "\n")

	for _, rowData := range rowsData {
		objectData := strings.Split(rowData, ",")

		if len(objectData) > 0 {
			ticketCSV := Ticket{
				Id: objectData[0],
				Nombre: objectData[1],
				Email: objectData[2],
				PaisDeDestino: objectData[3],
				HoraVuelo: objectData[4],
				Precio: objectData[5],
			}
			tickets = append(tickets, ticketCSV)
		}
	}

	return Tickets{Tickets: tickets}, nil
}

func main() {
	data, err := readFile("./tickets.csv")
	if err != nil {
		fmt.Println("Error en la lectura del archivo")
	}

	// resp, _:= data.GetTotalTickets("China")
	// fmt.Print(resp)

	// resp2, _:= data.GetMornings("noche")
	// fmt.Print(resp2)

	// resp3, _:= data.GetMornings("mañana")
	// fmt.Print(resp3)

	resp3,_:= data.AverageDestination("China")
	fmt.Print(resp3)
}
