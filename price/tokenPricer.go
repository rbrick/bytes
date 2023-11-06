package price

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rbrick/bytes/contracts"
	"github.com/rbrick/bytes/util"
)

const (
	OffchainOracle = "0x0AdDd25a91563696D8567Df78D5A01C9a991F9B8"
)

var (
	EtherToken = common.HexToAddress("0x0")

	EtherConverter PriceConverter = func(i *big.Int) *big.Float {
		// Ether has 18 decimal places
		return util.ConvertHuman(i, 18)
	}

	USDConverter PriceConverter = func(i *big.Int) *big.Float {
		// USDC has 6 decimal places
		return util.ConvertHuman(i, 6)
	}
)

type TokenPricer struct {
	rpc                 *ethclient.Client
	tokenAddress        common.Address
	denominationAddress common.Address
}

func (pricer *TokenPricer) Price() (*big.Int, error) {
	// This creates a new offchain oracle instance with the address to the oracle contract
	// and the given RPC for the token pricer.
	oracle, err := contracts.NewOffchainOracle(common.HexToAddress(OffchainOracle), pricer.rpc)

	if err != nil {
		return nil, err
	}

	// If the price is denominated in ether (the nil address), then we can call a dedicated function
	// on the contract 'getRateToEth' to get the price of the requested token in Ether
	if pricer.denominationAddress == EtherToken {
		return oracle.GetRateToEth(nil, pricer.tokenAddress, false)
	}

	// Else, the price is in another denomination (such as USDC), and should fetch the price accordingly.
	return oracle.GetRate(nil, pricer.tokenAddress, pricer.denominationAddress, false)
}

// NewTokenPricer creates a new token pricer.
func NewTokenPricer(rpc *ethclient.Client, tokenAddress, denominationAddress common.Address) Pricer {
	return &TokenPricer{
		rpc:                 rpc,
		tokenAddress:        tokenAddress,
		denominationAddress: denominationAddress,
	}
}
