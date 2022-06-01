package util

import (
	"github.com/holiman/uint256"
)

type SalaryPerBlock struct {
	AtAdvantage85  float64
	AtAdvantage100 float64
}

type SalaryConfig struct {
	MiningRatio        *uint256.Int
	BaseRewardPerBlock *uint256.Int
	K                  *uint256.Int
	RewardMultiplier   []*uint256.Int
	BSCDMultiplier     []*uint256.Int
}

type SalaryConfigData struct {
	MiningRatio        string   `json:"_miningRatio"`
	BaseRewardPerBlock string   `json:"_baseRewardPerBlock"`
	K                  string   `json:"_k"`
	RewardMultiplier   []string `json:"_rewardMultiplier"`
	BSCDMultiplier     []string `json:"_bscdMultiplier"`
}

type RefillFeeConfig struct {
	Refill80Fee []*uint256.Int
	Refill5Fee  []*uint256.Int
	Refill0Fee  []*uint256.Int
}

type RefillFeeConfigData struct {
	Refill80Fee []string `json:"_refill80Fee"`
	Refill5Fee  []string `json:"_refill5Fee"`
	Refill0Fee  []string `json:"_refill0Fee"`
}

type PvETicketFeeConfig struct {
	FeeBSCD []*uint256.Int
}

type PvETicketFeeConfigData struct {
	FeeBSCD []string `json:"_feeBSCD"`
}

type PillFeeItemConfig struct {
	AmountMST  *uint256.Int
	AmountBSCD *uint256.Int
}

type PillFeeConfig struct {
	Common    []PillFeeItemConfig
	Rare      []PillFeeItemConfig
	Elite     []PillFeeItemConfig
	Epic      []PillFeeItemConfig
	Legendary []PillFeeItemConfig
}

type PillFeeItemConfigData struct {
	AmountMST  string `json:"_amountMST"`
	AmountBSCD string `json:"_amountBSCD"`
}
type PillFeeConfigData struct {
	Common    []PillFeeItemConfigData `json:"_common"`
	Rare      []PillFeeItemConfigData `json:"_rare"`
	Elite     []PillFeeItemConfigData `json:"_elite"`
	Epic      []PillFeeItemConfigData `json:"_epic"`
	Legendary []PillFeeItemConfigData `json:"_legendary"`
}

type PillFeeRarityMultiplierConfig struct {
	Common    float64
	Rare      float64
	Elite     float64
	Epic      float64
	Legendary float64
}

type PillFeeRarityMultiplierConfigData struct {
	Common    string `json:"_common"`
	Rare      string `json:"_rare"`
	Elite     string `json:"_elite"`
	Epic      string `json:"_epic"`
	Legendary string `json:"_legendary"`
}

type PillFeeSystemValueConfigData struct {
	Value string `json:"_value"`
}

type PillFeeCurrencyPercentValueConfigData struct {
	MSTPercent  string `json:"_mstPercent"`
	BSCDPercent string `json:"_bscdPercent"`
}

type PillFeeCurrencyPercentValueConfig struct {
	MSTPercent  float64
	BSCDPercent float64
}
