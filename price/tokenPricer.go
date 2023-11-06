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
	oracle, err := contracts.NewOffchainOracle(common.HexToAddress(OffchainOracle), pricer.rpc)

	if err != nil {
		return nil, err
	}

	if pricer.denominationAddress == EtherToken {
		return oracle.GetRateToEth(nil, pricer.tokenAddress, false)
	}

	return oracle.GetRate(nil, pricer.tokenAddress, pricer.denominationAddress, false)
}
