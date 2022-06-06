package repo

import (
	"github.com/MStation-io/api/model"
)

type Repo struct {
	Transaction    TransactionRepo
	ProcessedBlock ProcessedBlockRepo
	User           UserRepo
	Class          ClassRepo
	Client         ClientRepo
	Course         CourseRepo
	CourseAtClass  CourseAtClassRepo
}

type TransactionRepo interface {
	Create(s DBRepo, param *model.Transaction) (*model.Transaction, error)
}

type ProcessedBlockRepo interface {
	Create(s DBRepo, param model.ProcessedBlock) (*model.ProcessedBlock, error)
	GetCurrentProcessedBlock(s DBRepo) (*model.ProcessedBlock, error)
	UpdateBlockNumberByBlockNumber(s DBRepo, blockNumber int64, updatedBlock int64) (*model.ProcessedBlock, error)
}

type UserRepo interface {
	Create(s DBRepo, param model.User) (*model.User, error)
	GetByReferralCode(s DBRepo, referralCode string) (*model.User, error)
	GetByWalletAddress(s DBRepo, walletAddress string) (*model.User, error)
	UpdateById(s DBRepo, id string, updateModel model.User) (*model.User, error)
	GetByWalletAddressWithReferrer(s DBRepo, walletAddress string) (*model.User, error)
	GetTotalFriendById(s DBRepo, ownerAddress string) (int64, error)
	GetFriendsWithReferralReward(s DBRepo, id string, address string) ([]model.User, error)
}

type ClassRepo interface {
	GetAllClasses(s DBRepo) ([]model.Class, error)
	PostClass(s DBRepo, param model.Class) (*model.Class, error)
}

type ClientRepo interface{}

type CourseRepo interface{}

type CourseAtClassRepo interface{}
