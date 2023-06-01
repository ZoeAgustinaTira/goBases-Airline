package tickets

import (
	"fmt"

	"github.com/ZoeAgustinaTira/goBases-Airline/internal/tickets/repository"
)

func GetTotalTicketsByDestination(destination string) (int, error) {
	count := 0
	for _, ticket := range repository.TicketList {
		if ticket.Country == destination {
			count++
		}
	}
	if count == 0 {
		return 0, fmt.Errorf("no passengers for destination: %s", destination)
	}

	return count, nil
}

/*
// ejemplo 2
func GetMornings(time string) (int, error) {}

// ejemplo 3
func AverageDestination(destination string, total int) (int, error) {}
*/
