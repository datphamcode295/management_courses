package pg

import (
	"github.com/MStation-io/api/model"
	"github.com/MStation-io/api/repo"
)

type classRepo struct{}

// func NewClassRepo() repo.ClassRepo {
// 	return &classRepo{}
// }

func GetAllClasses(s repo.DBRepo) ([]model.Class, error) { //

	tasks := []model.Class{}
	db := s.DB()

	return tasks, db.Find(&tasks).Error
}

func (t *classRepo) PostClass(s repo.DBRepo, param model.Class) (*model.Class, error) {

	db := s.DB()
	o := model.Class{ID: "4"}
	// return &param, db.Create(&param).Error

	x := db.Create(&o)
	return &o, x.Error

}

// func (t *classRepo) GetByReferralCode(s repo.DBRepo, referralCode string) (*model.Class, error) {
// 	db := s.DB()
// 	class := model.Class{}
// 	return &class, db.Where("referral_code = ?", referralCode).First(&class).Error
// }

// func (t *classRepo) GetByWalletAddress(s repo.DBRepo, walletAddress string) (*model.User, error) {
// 	db := s.DB()
// 	user := model.User{}
// 	return &user, db.Where("wallet_address ~~* ?", walletAddress).First(&user).Error
// }

// func (t *classRepo) GetByWalletAddressWithReferrer(s repo.DBRepo, walletAddress string) (*model.User, error) {
// 	db := s.DB()
// 	user := model.User{}
// 	return &user, db.Where("wallet_address ~~* ?", walletAddress).Preload("Referrer").First(&user).Error
// }

// func (t *classRepo) UpdateById(s repo.DBRepo, id string, updateModel model.User) (*model.User, error) {
// 	db := s.DB()
// 	user := model.User{}
// 	return &user, db.Model(&user).Where("id = ?", id).Updates(updateModel).Error
// }

// func (t *classRepo) GetTotalFriendById(s repo.DBRepo, id string) (int64, error) {
// 	db := s.DB()
// 	var total int64
// 	return total, db.Table("user").
// 		Where("referred_by = ?", id).
// 		Count(&total).Error
// }

// func (t *classRepo) GetFriendsWithReferralReward(s repo.DBRepo, id string, address string) ([]model.User, error) {
// 	db := s.DB()
// 	users := []model.User{}
// 	return users, db.
// 		Where("referred_by = ?", id).
// 		Preload("ReferredRewards", "owner_address ~~* ? AND status = ?", address, "claimed").
// 		Find(&users).Error
// }
