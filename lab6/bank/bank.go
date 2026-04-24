package bank

import "errors"

var (
	ErrInvalidAmount         = errors.New("amount must be greater than zero")
	ErrClientNotFound        = errors.New("client not found")
	ErrInsufficientFunds     = errors.New("client has insufficient funds")
	ErrInsufficientBankFunds = errors.New("bank has insufficient funds")
)

type Bank struct {
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
	return b.name
}

func (b *Bank) SetName(name string) {
	b.name = name
}

func (b *Bank) GetBankMoney() float64 {
	return b.bankMoney
}

func (b *Bank) SetBankMoney(bankMoney float64) {
	b.bankMoney = bankMoney
}

func (b *Bank) GetDeposit() float64 {
	return b.deposit
}

func (b *Bank) SetDeposit(deposit float64) {
	b.deposit = deposit
}

func (b *Bank) GetCredit() float64 {
	return b.credit
}

func (b *Bank) SetCredit(credit float64) {
	b.credit = credit
}

// AddClient stores a client by account number.
func (b *Bank) AddClient(client *Client) {
	if client == nil {
		return
	}

	b.clients[client.GetAccountNumber()] = client
}

func (b *Bank) GetClient(accountNumber string) (*Client, bool) {
	client, ok := b.clients[accountNumber]
	return client, ok
}

func (b *Bank) RemoveClient(accountNumber string) {
	delete(b.clients, accountNumber)
}

// GetClients returns a copy to keep internal storage encapsulated.
func (b *Bank) GetClients() map[string]*Client {
	clientsCopy := make(map[string]*Client, len(b.clients))
	for account, client := range b.clients {
		clientsCopy[account] = client
	}
	return clientsCopy
}

// SetClients replaces the registered clients using a defensive copy.
func (b *Bank) SetClients(clients map[string]*Client) {
	b.clients = make(map[string]*Client, len(clients))
	for account, client := range clients {
		b.clients[account] = client
	}
}

func (b *Bank) Deposit(accountNumber string, amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

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

// IssueCredit gives a client credit if the bank has enough available money.
func (b *Bank) IssueCredit(accountNumber string, amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

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
