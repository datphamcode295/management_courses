package pg

import (
	"github.com/MStation-io/api/model"
	"github.com/MStation-io/api/repo"
)

type courseAtClassRepo struct{}

func NewCourseAtClassRepo() repo.CourseAtClassRepo {
	return &courseAtClassRepo{}
}

func (t *courseAtClassRepo) GetAll(s repo.DBRepo) ([]model.CourseAtClass, error) {
	ats := []model.CourseAtClass{}
	db := s.DB()

	return ats, db.Find(&ats).Error
}

func (t *courseAtClassRepo) PostAts(s repo.DBRepo, param model.CourseAtClass) (*model.CourseAtClass, error) {
	db := s.DB()
	return &param, db.Create(&param).Error
}

func (t *courseAtClassRepo) DeleteByID(s repo.DBRepo, id string) (*model.CourseAtClass, error) {
	db := s.DB()
	var param model.CourseAtClass

	return &param, db.Where("id = ?", id).Delete(&param).Error
}

func (t *courseAtClassRepo) GetByID(s repo.DBRepo, id string) (*model.CourseAtClass, error) {
	db := s.DB()
	var class model.CourseAtClass

	return &class, db.Find(&class, "id=?", id).Error
}

func (t *courseAtClassRepo) GetByCourseID(s repo.DBRepo, id string) ([]model.CourseAtClass, error) {
	db := s.DB()
	var class []model.CourseAtClass

	return class, db.Find(&class, "course_id=?", id).Error
}

func (t *courseAtClassRepo) GetByClassID(s repo.DBRepo, id string) ([]model.CourseAtClass, error) {
	db := s.DB()
	var class []model.CourseAtClass

	return class, db.Find(&class, "class_id=?", id).Error
}
