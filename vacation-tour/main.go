package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"vacation-tour.com/pricing"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Vacation Tour Calculator")
	fmt.Println("Destinations: Bulgaria, Germany, Poland")
	fmt.Println("Seasons: summer, winter")
	fmt.Print("Destination: ")
	destinationRaw, _ := reader.ReadString('\n')
	destination := strings.TrimSpace(destinationRaw)

	fmt.Print("Season: ")
	seasonRaw, _ := reader.ReadString('\n')
	season := strings.TrimSpace(seasonRaw)

	fmt.Print("Days: ")
	daysRaw, _ := reader.ReadString('\n')
	days, err := strconv.Atoi(strings.TrimSpace(daysRaw))
	if err != nil {
		fmt.Println("Input error: please enter a valid integer for days")
		return
	}

	fmt.Print("Room type (standard/luxury): ")
	roomTypeRaw, _ := reader.ReadString('\n')
	roomType := strings.TrimSpace(roomTypeRaw)

	total, err := pricing.CalculateTourPrice(destination, season, days, roomType)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Total price: $%.2f\n", total)
}
