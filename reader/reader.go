package reader

import (
	"encoding/csv"
	"fmt"
	"github.com/ZoeAgustinaTira/goBases-Airline/domain"
	"github.com/ZoeAgustinaTira/goBases-Airline/internal/tickets/repository"
	"os"
)

func ReadFile() {
	file, err := os.Open("./../../tickets.csv")

	nr := csv.NewReader(file)
	data, err := nr.ReadAll()
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, row := range data {
		repository.TicketList = append(repository.TicketList, domain.Ticket{
			ID:      row[0],
			Name:    row[1],
			Mail:    row[2],
			Country: row[3],
			Time:    row[4],
			Amount:  row[5],
		})

	}
}
