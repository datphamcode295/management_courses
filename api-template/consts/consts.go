package consts

const (
	DefaultSignatureLengthWithRecoveryId = 65

	//Default Lock Time(in second)
	DefaultLockTime         = 420
	DefaultGetPriceInterval = 120
	DefaultContextTimeout   = 10
	DefaultVentureTimeout   = 600
	DefaultCronLogInterval  = 3600

	//Advisory locks
	AdvisoryLockTime                 = 600
	AdvisoryCronjobNamespace         = 0
	AdvisoryLockContractConfig       = 0
	AdvisoryLockResetAdventureChance = 1

	//Action name
	ActionCharacterMinting = "character_minted"
)
