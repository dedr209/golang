package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"lab6/bank"
)

type appState struct {
	mu    sync.RWMutex
	bots  map[string]map[string]chan struct{}
	banks map[string]*bank.Bank
	conf  bank.BotConfig
	wg    sync.WaitGroup
}

func newAppState() *appState {
	return &appState{
		bots:  make(map[string]map[string]chan struct{}),
		banks: make(map[string]*bank.Bank),
		conf:  bank.DefaultBotConfig(),
	}
}

func (a *appState) createBank(name string, money, deposit, credit float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if _, exists := a.banks[name]; exists {
		return fmt.Errorf("bank %q already exists", name)
	}

	a.banks[name] = bank.NewBank(name, money, deposit, credit)
	a.bots[name] = make(map[string]chan struct{})
	return nil
}

func (a *appState) createClientAndStartBot(bankName, name, surname, account string, deposit, credit float64) error {
	a.mu.RLock()
	b, exists := a.banks[bankName]
	_, botExists := a.bots[bankName][account]
	a.mu.RUnlock()
	if !exists {
		return fmt.Errorf("bank %q not found", bankName)
	}
	if botExists {
		return fmt.Errorf("account %q already has an active bot in bank %q", account, bankName)
	}

	client := bank.NewClient(name, surname, account, deposit, credit)
	b.AddClient(client)

	stop := make(chan struct{})
	a.mu.Lock()
	a.bots[bankName][account] = stop
	a.mu.Unlock()

	a.wg.Add(1)
	go client.RunBotWithStop(b, a.conf, &a.wg, stop)
	return nil
}

func (a *appState) printStatus() {
	a.mu.RLock()
	banks := make(map[string]*bank.Bank, len(a.banks))
	for name, b := range a.banks {
		banks[name] = b
	}
	a.mu.RUnlock()

	if len(banks) == 0 {
		fmt.Println("No banks created yet.")
		return
	}

	for bankName, b := range banks {
		fmt.Printf("\nBank: %s | Money: %.2f | Deposit: %.2f | Credit: %.2f\n", bankName, b.GetBankMoney(), b.GetDeposit(), b.GetCredit())
		clients := b.GetClients()
		if len(clients) == 0 {
			fmt.Println("  clients: none")
			continue
		}

		for account, client := range clients {
			dep, cre, _ := b.GetClientBalances(account)
			fmt.Printf("  %s %s (%s) -> deposit: %.2f, credit: %.2f\n", client.GetName(), client.GetSurname(), account, dep, cre)
		}
	}
}

func (a *appState) stopAllBots() {
	a.mu.Lock()
	for _, byAccount := range a.bots {
		for account, stop := range byAccount {
			close(stop)
			delete(byAccount, account)
		}
	}
	a.mu.Unlock()

	a.wg.Wait()
}

func readLine(reader *bufio.Reader, prompt string) string {
	for {
		fmt.Print(prompt)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("input error, try again")
			continue
		}
		value := strings.TrimSpace(line)
		if value != "" {
			return value
		}
		fmt.Println("value cannot be empty")
	}
}

func readFloat(reader *bufio.Reader, prompt string) float64 {
	for {
		text := readLine(reader, prompt)
		value, err := strconv.ParseFloat(text, 64)
		if err != nil {
			fmt.Println("enter a valid number")
			continue
		}
		return value
	}
}

func main() {
	state := newAppState()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Bank Menu ---")
		fmt.Println("1) Create bank")
		fmt.Println("2) Create client and start bot")
		fmt.Println("3) View status")
		fmt.Println("4) Exit")

		choice := readLine(reader, "Choose action: ")
		switch choice {
		case "1":
			name := readLine(reader, "Bank name: ")
			money := readFloat(reader, "Initial bank money: ")
			deposit := readFloat(reader, "Initial total deposit: ")
			credit := readFloat(reader, "Initial total credit: ")
			if err := state.createBank(name, money, deposit, credit); err != nil {
				fmt.Printf("error: %v\n", err)
				continue
			}
			fmt.Println("bank created")
		case "2":
			bankName := readLine(reader, "Bank name: ")
			name := readLine(reader, "Client name: ")
			surname := readLine(reader, "Client surname: ")
			account := readLine(reader, "Account number: ")
			deposit := readFloat(reader, "Client deposit: ")
			credit := readFloat(reader, "Client credit: ")

			if err := state.createClientAndStartBot(bankName, name, surname, account, deposit, credit); err != nil {
				fmt.Printf("error: %v\n", err)
				continue
			}
			fmt.Println("client created and bot started")
		case "3":
			state.printStatus()
		case "4":
			fmt.Println("Stopping bots...")
			state.stopAllBots()
			fmt.Println("Goodbye.")
			return
		default:
			fmt.Println("unknown option")
		}
	}
}
