package main

import (
	"fmt"
	"sync"

	"lab6/bank"
)

func main() {
	myBank := bank.NewBank("BVO Bank", 1500000, 350000, 120000)
	clients := []*bank.Client{
		bank.NewClient("Ivan", "Petrenko", "UA-1001", 5000, 1200),
		bank.NewClient("Olena", "Melnyk", "UA-1002", 2500, 600),
		bank.NewClient("Maksym", "Bondar", "UA-1003", 1500, 300),
	}

	for _, client := range clients {
		myBank.AddClient(client)
	}

	fmt.Printf("Bank: %s | Money: %.2f\n", myBank.GetName(), myBank.GetBankMoney())

	config := bank.DefaultBotConfig()
	var wg sync.WaitGroup
	for _, client := range clients {
		wg.Add(1)
		go client.RunBot(myBank, config, &wg)
	}

	wg.Wait()
	fmt.Println("All client bots finished their work.")
}
