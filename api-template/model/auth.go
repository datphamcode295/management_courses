package model

// ChallengeInfo ...
type ChallengeInfo struct {
	ExpiresAt int64  `json:"exp"`
	Address   string `json:"address"`
}
