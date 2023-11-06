package wallet

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rbrick/bytes/contracts"
	"github.com/rbrick/bytes/util"
)

var (
	NeoTokyoStakerAddress = common.HexToAddress("0x67e1eCFA9232E27EAf3133B968A33A9a0dCa9e16")
)

type Position struct {
	Season        int        `json:"season"`
	Id            int        `json:"id"`
	Points        int        `json:"points"`
	StakedBytes   *big.Float `json:"stakedBytes"`
	StakedVaultId string     `json:"stakedVaultId"`
	TimeLockEnd   int64      `json:"timelockEnd"`
	TimeLock      int64      `json:"timelock"`
}

func GetPositions(stakerContract *contracts.NeoTokyoStaker, address common.Address) ([]*Position, error) {
	stakerPositions, err := stakerContract.GetStakerPositions(nil, address)

	if err != nil {
		return nil, err
	}

	var positions []*Position

	for _, position := range stakerPositions.StakedS1Citizens {
		positions = append(positions, &Position{
			Season:        1,
			Id:            int(position.CitizenId.Int64()),
			Points:        int(position.Points.Int64()),
			StakedBytes:   util.ConvertHuman(position.StakedBytes, 18),
			StakedVaultId: position.StakedVaultId.String(),
			TimeLockEnd:   position.TimelockEndTime.Int64(),
		})
	}

	for _, position := range stakerPositions.StakedS2Citizens {
		positions = append(positions, &Position{
			Season:        2,
			Id:            int(position.CitizenId.Int64()),
			Points:        int(position.Points.Int64()),
			StakedBytes:   util.ConvertHuman(position.StakedBytes, 18),
			StakedVaultId: "N/A",
			TimeLockEnd:   position.TimelockEndTime.Int64(),
		})
	}

	return positions, nil
}
