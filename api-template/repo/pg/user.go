package pg

import (
	"github.com/MStation-io/api/model"
	"github.com/MStation-io/api/repo"
)

type userRepo struct{}

func NewUserRepo() repo.UserRepo {
	return &userRepo{}
}

func (t *userRepo) Create(s repo.DBRepo, param model.User) (*model.User, error) {
	db := s.DB()
	return &param, db.Create(&param).Error
}

func (t *userRepo) GetByReferralCode(s repo.DBRepo, referralCode string) (*model.User, error) {
	db := s.DB()
	user := model.User{}
	return &user, db.Where("referral_code = ?", referralCode).First(&user).Error
}

func (t *userRepo) GetByWalletAddress(s repo.DBRepo, walletAddress string) (*model.User, error) {
	db := s.DB()
	user := model.User{}
	return &user, db.Where("wallet_address ~~* ?", walletAddress).First(&user).Error
}

func (t *userRepo) GetByWalletAddressWithReferrer(s repo.DBRepo, walletAddress string) (*model.User, error) {
	db := s.DB()
	user := model.User{}
	return &user, db.Where("wallet_address ~~* ?", walletAddress).Preload("Referrer").First(&user).Error
}

func (t *userRepo) UpdateById(s repo.DBRepo, id string, updateModel model.User) (*model.User, error) {
	db := s.DB()
	user := model.User{}
	return &user, db.Model(&user).Where("id = ?", id).Updates(updateModel).Error
}

func (t *userRepo) GetTotalFriendById(s repo.DBRepo, id string) (int64, error) {
	db := s.DB()
	var total int64
	return total, db.Table("user").
		Where("referred_by = ?", id).
		Count(&total).Error
}

func (t *userRepo) GetFriendsWithReferralReward(s repo.DBRepo, id string, address string) ([]model.User, error) {
	db := s.DB()
	users := []model.User{}
	return users, db.
		Where("referred_by = ?", id).
		Preload("ReferredRewards", "owner_address ~~* ? AND status = ?", address, "claimed").
		Find(&users).Error
}
