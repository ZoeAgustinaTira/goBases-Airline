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
