package blockchain

type BlockChainValidator interface {
	ValidateTxHashMintCharacter(txHash string, minter string) error
	ValidateTxHashTransferCharacter(txHash string, sender string, nftAddress string) error
	ValidateTxHashClaimReferralReward(txHash string, ownerAddress string, amount int64, claimId int64) error
	ValidateTxHashClaimSalary(txHash string, ownerAddress string) error
	ValidateTxHashQuitWorking(txHash string, ownerAddress string) error
	ValidateTxHashStartWorking(txHash string, ownerAddress string) error
	ValidateTxHashUpgradeCharacter(txHash string, ownerAddress string) error
	ValidateTxHashRefillEnergy(txHash string, ownerAddress string) error
	ValidateTxHashStartAdventure(txHash string, ownerAddress string) error
	ValidateTxHashMintPackCharacter(txHash string, minter string) error
}
