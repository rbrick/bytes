package models

type Prices struct {
	USDPrice   string `json:"usdPrice"`
	EtherPrice string `json:"etherPrice"`
}

type Wallet struct {
	// The balance of the 'BYTES' token
	TokenBalance string `json:"balance"`
	// The total amount of tokens staked
	TotalStake string `json:"totalStake"`
	// The amount of 'BYTES' tokens yielded.
	PendingRewards string `json:"pendingRewards"`

	// The citizens staked for this wallet
	StakedCitizens []Citizen `json:"stakedCitizens"`
}

type Citizen struct {
	// The ID of this citizen
	ID int `json:"id"`
	// The season of this citizen
	Season int `json:"season"`
	// The image for this citizen
	Image string `json:"image"`
}
