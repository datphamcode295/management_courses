package pg

import "github.com/MStation-io/api/repo"

func NewRepo() *repo.Repo {
	return &repo.Repo{
		Transaction:    NewTransactionRepo(),
		ProcessedBlock: NewProcessedBlockRepoo(),
		User:           NewUserRepo(),
		Class:          NewClassRepo(),
		Course:         NewCourseRepo(),
		CourseAtClass:  NewCourseAtClassRepo(),
		Client:         NewClientRepo(),
	}
}
