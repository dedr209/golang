package bank

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Client stores personal and account-level financial data.
type Client struct {
	name          string
	surname       string
	accountNumber string
	cDeposit      float64
	cCredit       float64
}

func NewClient(name, surname, accountNumber string, cDeposit, cCredit float64) *Client {
	return &Client{
		name:          name,
		surname:       surname,
		accountNumber: accountNumber,
		cDeposit:      cDeposit,
		cCredit:       cCredit,
	}
}

func (c *Client) GetName() string {
	return c.name
}

func (c *Client) SetName(name string) {
	c.name = name
}

func (c *Client) GetSurname() string {
	return c.surname
}

func (c *Client) SetSurname(surname string) {
	c.surname = surname
}

func (c *Client) GetAccountNumber() string {
	return c.accountNumber
}

func (c *Client) SetAccountNumber(accountNumber string) {
	c.accountNumber = accountNumber
}

func (c *Client) GetDeposit() float64 {
	return c.cDeposit
}

func (c *Client) SetDeposit(cDeposit float64) {
	c.cDeposit = cDeposit
}

func (c *Client) GetCredit() float64 {
	return c.cCredit
}

func (c *Client) SetCredit(cCredit float64) {
	c.cCredit = cCredit
}

// BotConfig controls how frequently and how aggressively a client bot acts.
type BotConfig struct {
	Interval       time.Duration
	MinAmount      float64
	MaxAmount      float64
	WithdrawChance int
}

// DefaultBotConfig runs an action every second.
func DefaultBotConfig() BotConfig {
	return BotConfig{
		Interval:       time.Second,
		MinAmount:      10,
		MaxAmount:      100,
		WithdrawChance: 70,
	}
}

func normalizeBotConfig(config BotConfig) BotConfig {
	if config.Interval <= 0 {
		config.Interval = time.Second
	}
	if config.MinAmount <= 0 {
		config.MinAmount = 10
	}
	if config.MaxAmount < config.MinAmount {
		config.MaxAmount = config.MinAmount
	}
	if config.WithdrawChance < 0 {
		config.WithdrawChance = 0
	}
	if config.WithdrawChance > 100 {
		config.WithdrawChance = 100
	}

	return config
}

// RunBot starts an infinite client loop and should be launched with the go keyword.
func (c *Client) RunBot(b *Bank, config BotConfig, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}

	config = normalizeBotConfig(config)
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	accountNumber := c.GetAccountNumber()
	initialCredit := c.GetCredit()

	for {
		time.Sleep(config.Interval)

		amount := config.MinAmount
		if config.MaxAmount > config.MinAmount {
			amount += rng.Float64() * (config.MaxAmount - config.MinAmount)
		}

		if rng.Intn(100) < config.WithdrawChance {
			_ = b.Withdraw(accountNumber, amount)
		} else {
			_ = b.Deposit(accountNumber, amount)
		}

		if _, credit, ok := b.GetClientBalances(accountNumber); ok && credit > 0 {
			payAmount := amount / 2
			if payAmount > credit {
				payAmount = credit
			}
			_ = b.PayCredit(accountNumber, payAmount)
		}

		deposit, credit, ok := b.GetClientBalances(accountNumber)
		if !ok {
			fmt.Printf("client %s removed from bank, bot stopped\n", accountNumber)
			return
		}

		if deposit <= 0 {
			fmt.Printf("client %s reached empty deposit, bot stopped\n", accountNumber)
			return
		}
		if initialCredit > 0 && credit <= 0 {
			fmt.Printf("client %s paid off credit, bot stopped\n", accountNumber)
			return
		}
	}
}
