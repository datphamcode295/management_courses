package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//User ..
type User struct {
	Id            string     `gorm:"PRIMARY_KEY;" json:"Id"`
	WalletAddress string     `json:"wallet_address"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	ReferralCode  string     `json:"referral_code"`
	ReferredBy    string     `json:"referred_by"`
	ReferralDate  *time.Time `json:"referral_date"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()
	u.ReferralCode, _ = GenerateUniqueNanoId(15)
	return nil
}

func (User) TableName() string {
	return "user"
}
