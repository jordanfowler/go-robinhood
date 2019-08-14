package robinhood

// Orderbook is a representation of the data returned by the Robinhood Gold API for
// Nasdaq Level II bid/ask order book data
type Orderbook struct {
	Symbol string  `json:"symbol"`
	Asks   []Order `json:"asks"`
	Bids   []Order `json:"bids"`
}

// Order represents a single order in an Orderbook
type Order struct {
	Side     string `json:"side"`
	Price    Price  `json:"price"`
	Quantity uint64 `json:"quantity"`
}

// Price is the Order price
type Price struct {
	Amount       string `json:"amount"`
	CurrencyCode string `json:"currency_code"`
}

// GetOrderbook returns the latest order book for the given ticker symbol
func (c *Client) GetOrderbook(sym string) (*Orderbook, error) {
	instr, err := c.GetInstrumentForSymbol(sym)
	if err != nil {
		return nil, err
	}
	url := EPMarketbook + instr.ID
	var r Orderbook
	err = c.GetAndDecode(url, &r)
	r.Symbol = sym
	return &r, err
}
