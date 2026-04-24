package bank

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
