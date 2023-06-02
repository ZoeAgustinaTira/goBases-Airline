package tickets

import (
	"fmt"
	"github.com/ZoeAgustinaTira/goBases-Airline/domain"
	"github.com/ZoeAgustinaTira/goBases-Airline/internal/tickets/repository"
	"strconv"
	"strings"
)

func AverageDestination(ch chan string, destination string) {
	var ticketbc []domain.Ticket

	for _, ticket := range repository.TicketList {
		if ticket.Country == destination {
			ticketbc = append(ticketbc, ticket)
		}
	}
	avg := float64(len(ticketbc)) * 100.0 / float64(len(repository.TicketList))

	ch <- fmt.Sprintf("The average number of flights to the destination %s is %.2f", destination, avg)

}

func GetTotalTicketsByDestination(ch chan string, destination string) {
	count := 0
	for _, ticket := range repository.TicketList {
		if ticket.Country == destination {
			count++
		}
	}
	if count == 0 {
		ch <- fmt.Sprintf("no passengers for destination: %s", destination)
	}

	ch <- fmt.Sprintf("There are %d passengers flying to %s", count, destination)
}

func GetCountByPeriod(ch chan string, period string) {
	var from, to int
	switch period {
	case "earlyMorning":
		from, to = 0, 6
	case "morning":
		from, to = 7, 12
	case "evening":
		from, to = 13, 19
	case "night":
		from, to = 20, 23
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error during processing:", err)
		}
	}()

	count := 0
	for _, ticket := range repository.TicketList {
		hour, err := strconv.Atoi(strings.Split(ticket.Time, ":")[0])
		if err != nil {
			panic("could not split")
		}
		if hour >= from && hour <= to {
			count++
		}
	}

	ch <- fmt.Sprintf("There are %d passengers flying in that time period %s", count, period)

}
