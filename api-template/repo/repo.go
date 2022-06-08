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
	GetClassByID(s DBRepo, id string) (*model.Class, error)
	GetClassByName(s DBRepo, id string) (*model.Class, error)
	GetAllStudent(s DBRepo, name string) ([]model.Client, error)
	DeleteByID(s DBRepo, id string) (*model.Class, error)
}

type ClientRepo interface {
	GetAllClients(s DBRepo) ([]model.Client, error)
	PostClient(s DBRepo, param model.Client) (*model.Client, error)
	Delete(s DBRepo, id string) (*model.Client, error)
	UpdatePassword(s DBRepo, id string, password string) (*model.Client, error)
	CheckingLogin(s DBRepo, username string, password string) (*model.Client, error)
	GetByUsername(s DBRepo, username string) (*model.Client, error)
	GetByID(s DBRepo, id string) (*model.Client, error)
}

type CourseRepo interface {
	GetAllCourses(s DBRepo) ([]model.Course, error)
	PostCourse(s DBRepo, param model.Course) (*model.Course, error)
	DeleteByID(s DBRepo, id string) (*model.Course, error)
	GetCourseById(s DBRepo, id string) (*model.Course, error)
}
type CourseAtClassRepo interface {
	GetAll(s DBRepo) ([]model.CourseAtClass, error)
	PostAts(s DBRepo, param model.CourseAtClass) (*model.CourseAtClass, error)
	DeleteByID(s DBRepo, id string) (*model.CourseAtClass, error)
	GetByID(s DBRepo, id string) (*model.CourseAtClass, error)
	GetByCourseID(s DBRepo, id string) ([]model.CourseAtClass, error)
	GetByClassID(s DBRepo, id string) ([]model.CourseAtClass, error)
}
