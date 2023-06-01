package tickets

import (
	"github.com/ZoeAgustinaTira/goBases-Airline/domain"
	"github.com/ZoeAgustinaTira/goBases-Airline/internal/tickets/repository"
)

/*
// ejemplo 1
func GetTotalTickets(destination string) (int, error) {

		return 0,nil
	}

// ejemplo 2
func GetMornings(time string) (int, error) {}
*/
func AverageDestination(destination string) (float64, error) {
	var ticketbc []domain.Ticket

	for _, ticket := range repository.TicketList {
		if ticket.Country == destination {
			ticketbc = append(ticketbc, ticket)
		}
	}
	avg := float64(len(ticketbc)) * 100.0 / float64(len(repository.TicketList))

	return avg, nil
}
