package util

import (
	"math"
	"math/big"
)

const (
	Gwei  = 1000000000
	Ether = 1000000000000000000
)

const (
	Wei = 1e18
)

func ConvertEtherUnits(amount *big.Float, decimalPlaces int) *big.Int {
	divisor := math.Pow10(decimalPlaces)

	var integer *big.Int

	if Wei/divisor == 1 {
		integer, _ = new(big.Float).Mul(amount, big.NewFloat(Wei)).Int(nil)
	} else {
		integer, _ = new(big.Float).Mul(amount, big.NewFloat(Wei/divisor)).Int(nil)
	}
	return integer
}

func ConvertCurrency(amount *big.Float, decimalPlaces int) *big.Int {
	multiplier := math.Pow10(decimalPlaces)
	integer, _ := new(big.Float).Mul(amount, big.NewFloat(multiplier)).Int(nil)
	return integer
}

func ConvertHuman(amount *big.Int, decimalPlaces int) *big.Float {
	divisor := math.Pow10(decimalPlaces)
	floatAmount := new(big.Float).SetInt(amount)
	return new(big.Float).Quo(floatAmount, big.NewFloat(divisor))
}
