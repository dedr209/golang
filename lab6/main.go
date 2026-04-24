package main

import (
	"fmt"

	"lab6/bank"
)

func main() {
	myBank := bank.NewBank("BVO Bank", 1500000, 350000, 120000)
	client := bank.NewClient("Ivan", "Petrenko", "UA-1001", 5000, 1200)

	myBank.AddClient(client)

	fmt.Printf("Bank: %s | Money: %.2f\n", myBank.GetName(), myBank.GetBankMoney())
	if storedClient, ok := myBank.GetClient("UA-1001"); ok {
		fmt.Printf("Client: %s %s | Deposit: %.2f | Credit: %.2f\n",
			storedClient.GetName(),
			storedClient.GetSurname(),
			storedClient.GetDeposit(),
			storedClient.GetCredit(),
		)
	}
}
