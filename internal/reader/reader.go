package reader

import (
	"encoding/csv"
	"github.com/ZoeAgustinaTira/goBases-Airline/internal/domain"
	"github.com/ZoeAgustinaTira/goBases-Airline/internal/repository"
	"os"
)

func ReadFile() {
	//	var ticketList []domain.Ticket
	//rawData, err := os.ReadFile("./../../tickets.csv")
	file, err := os.Open("./../../tickets.csv")
	dc := csv.NewReader(file)
	data, err := dc.ReadAll()
	if err != nil {
		panic(err.Error())
	}
	/*
		err = json.Unmarshal(rawData, ticketList)
		if err != nil {
			panic(err)
		}
	*/
	//	fmt.Println(ticketList)

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

	//fmt.Println(ticketList[1])

}
