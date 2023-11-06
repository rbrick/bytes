package models

type Prices struct {
	USDPrice   string `json:"usdPrice"`
	EtherPrice string `json:"etherPrice"`
}

type Wallet struct {
	// The balance of the 'BYTES' token
	TokenBalance string `json:"balance"`
	// The amount of 'BYTES' tokens yielded.
	YieldedTokens string `json:"yield"`
	// The citizens staked for this wallet
	StakedCitizens []Citizen `json:"stakedCitizens"`
}

type Citizen struct {
	// The ID of this citizen
	ID int `json:"id"`
	// The image for this citizen
	Image string `json:"image"`
}
