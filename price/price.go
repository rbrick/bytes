package price

import (
	"math/big"
)

// A pricer fetches the price of a token
type Pricer interface {
	Price() (*big.Int, error)
}
