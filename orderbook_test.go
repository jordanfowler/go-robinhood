package robinhood

import (
	"fmt"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestOrderbook(t *testing.T) {
	if os.Getenv("ROBINHOOD_USERNAME") == "" {
		t.Skip("No username set")
		return
	}
	asrt := assert.New(t)
	o := &OAuth{
		Username: os.Getenv("ROBINHOOD_USERNAME"),
		Password: os.Getenv("ROBINHOOD_PASSWORD"),
	}

	c, err := Dial(&CredsCacher{Creds: o})

	asrt.NoError(err)
	asrt.NotNil(c)

	sym := "SPY"
	ob, err := c.GetOrderbook(sym)
	asrt.Equal(sym, ob.Symbol)
	asrt.NoError(err)

	spew.Dump(ob)
	fmt.Printf("ob.Asks = %+v ob.Bids = %+v\n", len(ob.Asks), len(ob.Bids))
}
