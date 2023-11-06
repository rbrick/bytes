package price

import "math/big"

// PriceConverter converts an integer into a human readable floating point number.
// In Ethereum, everything is handled through big integers, while this is ideal for computers
// this isn't the case for displaying the actual price.
type PriceConverter func(*big.Int) *big.Float
