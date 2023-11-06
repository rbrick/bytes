package wallet

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rbrick/bytes/contracts"
)

func GetPendingRewards(stakerContract *contracts.NeoTokyoStaker, address common.Address) (*big.Int, error) {
	// 0 is Season 1
	// 1 is Season 2
	rewardS1, _, err := stakerContract.GetPendingPoolReward(nil, 0, address)

	if err != nil {
		return nil, err
	}

	rewardS2, _, err := stakerContract.GetPendingPoolReward(nil, 1, address)

	if err != nil {
		return nil, err
	}

	return new(big.Int).Add(rewardS1, rewardS2), nil
}
