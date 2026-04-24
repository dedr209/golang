package bank

import (
	"errors"
	"sync"
)

var (
	ErrInvalidAmount         = errors.New("amount must be greater than zero")
	ErrClientNotFound        = errors.New("client not found")
	ErrInsufficientFunds     = errors.New("client has insufficient funds")
	ErrInsufficientBankFunds = errors.New("bank has insufficient funds")
)

type Bank struct {
	mu        sync.RWMutex
	name      string
	bankMoney float64
	deposit   float64
	credit    float64
	clients   map[string]*Client
}

// NewBank constructs a Bank with initial values.
func NewBank(name string, bankMoney, deposit, credit float64) *Bank {
	return &Bank{
		name:      name,
		bankMoney: bankMoney,
		deposit:   deposit,
		credit:    credit,
		clients:   make(map[string]*Client),
	}
}

func (b *Bank) GetName() string {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.name
}

func (b *Bank) SetName(name string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.name = name
}

func (b *Bank) GetBankMoney() float64 {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.bankMoney
}

func (b *Bank) SetBankMoney(bankMoney float64) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.bankMoney = bankMoney
}

func (b *Bank) GetDeposit() float64 {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.deposit
}

func (b *Bank) SetDeposit(deposit float64) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.deposit = deposit
}

func (b *Bank) GetCredit() float64 {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.credit
}

func (b *Bank) SetCredit(credit float64) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.credit = credit
}

func (b *Bank) AddClient(client *Client) {
	if client == nil {
		return
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	b.clients[client.GetAccountNumber()] = client
}

func (b *Bank) GetClient(accountNumber string) (*Client, bool) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	client, ok := b.clients[accountNumber]
	return client, ok
}

func (b *Bank) RemoveClient(accountNumber string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	delete(b.clients, accountNumber)
}

// GetClients returns a copy to keep internal storage encapsulated.
func (b *Bank) GetClients() map[string]*Client {
	b.mu.RLock()
	defer b.mu.RUnlock()

	clientsCopy := make(map[string]*Client, len(b.clients))
	for account, client := range b.clients {
		clientsCopy[account] = client
	}
	return clientsCopy
}

func (b *Bank) SetClients(clients map[string]*Client) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.clients = make(map[string]*Client, len(clients))
	for account, client := range clients {
		b.clients[account] = client
	}
}

func (b *Bank) Deposit(accountNumber string, amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	b.mu.Lock()
	defer b.mu.Unlock()

	client, ok := b.clients[accountNumber]
	if !ok {
		return ErrClientNotFound
	}

	client.cDeposit += amount
	b.deposit += amount
	b.bankMoney += amount

	return nil
}

func (b *Bank) Withdraw(accountNumber string, amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	b.mu.Lock()
	defer b.mu.Unlock()

	client, ok := b.clients[accountNumber]
	if !ok {
		return ErrClientNotFound
	}
	if client.cDeposit < amount {
		return ErrInsufficientFunds
	}
	if b.bankMoney < amount {
		return ErrInsufficientBankFunds
	}

	client.cDeposit -= amount
	b.deposit -= amount
	b.bankMoney -= amount

	return nil
}

func (b *Bank) IssueCredit(accountNumber string, amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	b.mu.Lock()
	defer b.mu.Unlock()

	client, ok := b.clients[accountNumber]
	if !ok {
		return ErrClientNotFound
	}
	if b.bankMoney < amount {
		return ErrInsufficientBankFunds
	}

	client.cCredit += amount
	client.cDeposit += amount
	b.credit += amount
	b.deposit += amount
	b.bankMoney -= amount

	return nil
}

func (b *Bank) PayCredit(accountNumber string, amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	b.mu.Lock()
	defer b.mu.Unlock()

	client, ok := b.clients[accountNumber]
	if !ok {
		return ErrClientNotFound
	}
	if client.cDeposit < amount || client.cCredit < amount {
		return ErrInsufficientFunds
	}

	client.cDeposit -= amount
	client.cCredit -= amount
	b.deposit -= amount
	b.credit -= amount
	b.bankMoney += amount

	return nil
}

func (b *Bank) GetClientBalances(accountNumber string) (float64, float64, bool) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	client, ok := b.clients[accountNumber]
	if !ok {
		return 0, 0, false
	}

	return client.cDeposit, client.cCredit, true
}
