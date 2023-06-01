package tickets

import (
	"fmt"
	"github.com/ZoeAgustinaTira/goBases-Airline/internal/tickets/repository"
	"strconv"
	"strings"
)

/*
// ejemplo 1
func GetTotalTickets(destination string) (int, error) {

	return 0,nil
}

// ejemplo 2
func GetMornings(time string) (int, error) {}

// ejemplo 3
func AverageDestination(destination string, total int) (int, error) {}
*/
func GetCountByPeriod(periodo string) {
	var inicio, fin int
	switch periodo {
	case "madrugada":
		inicio, fin = 0, 6
	case "mañana":
		inicio, fin = 7, 12
	case "tarde":
		inicio, fin = 13, 19
	case "noche":
		inicio, fin = 20, 23
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error durante el procesamiento:", err)
		}
	}()

	contar := 0
	for _, ticket := range repository.TicketList {
		hora, err := strconv.Atoi(strings.Split(ticket.Time, ":")[0])
		if err != nil {
			panic("No se pudo dividir el string de la hora")
		}
		if hora >= inicio && hora <= fin {
			contar++
		}
	}

	fmt.Printf("Hay %d pasajeros volando en el periodo '%s'", contar, periodo)
}
