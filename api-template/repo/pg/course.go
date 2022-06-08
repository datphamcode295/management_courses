package pg

import (
	"github.com/MStation-io/api/model"
	"github.com/MStation-io/api/repo"
)

type courseRepo struct{}

func NewCourseRepo() repo.CourseRepo {
	return &courseRepo{}
}

func (t *courseRepo) GetAllCourses(s repo.DBRepo) ([]model.Course, error) { //

	courses := []model.Course{}
	db := s.DB()

	return courses, db.Find(&courses).Error
}

func (t *courseRepo) GetCourseById(s repo.DBRepo, id string) (*model.Course, error) {
	course := model.Course{}
	db := s.DB()

	return &course, db.Find(&course, "id=?", id).Error
}

func (t *courseRepo) PostCourse(s repo.DBRepo, param model.Course) (*model.Course, error) {
	db := s.DB()
	return &param, db.Create(&param).Error

}

func (t *courseRepo) DeleteByID(s repo.DBRepo, id string) (*model.Course, error) {
	db := s.DB()
	var param model.Course

	return &param, db.Where("id = ?", id).Delete(&param).Error
}
