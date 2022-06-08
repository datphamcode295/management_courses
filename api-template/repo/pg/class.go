package pg

import (
	"github.com/MStation-io/api/model"
	"github.com/MStation-io/api/repo"
)

type classRepo struct{}

func NewClassRepo() repo.ClassRepo {
	return &classRepo{}
}

func (t *classRepo) GetAllClasses(s repo.DBRepo) ([]model.Class, error) { //

	tasks := []model.Class{}
	db := s.DB()

	return tasks, db.Find(&tasks).Error
}

func (t *classRepo) PostClass(s repo.DBRepo, param model.Class) (*model.Class, error) {

	db := s.DB()

	// return &param, db.Create(&param).Error

	x := db.Create(&param)
	return &param, x.Error

}

func (t *classRepo) GetClassByID(s repo.DBRepo, id string) (*model.Class, error) {
	db := s.DB()
	var class model.Class

	return &class, db.Find(&class, "id=?", id).Error
}

func (t *classRepo) GetClassByName(s repo.DBRepo, name string) (*model.Class, error) {
	db := s.DB()
	var class model.Class

	return &class, db.Find(&class, "name=?", name).Error
}

func (t *classRepo) GetAllStudent(s repo.DBRepo, id string) ([]model.Client, error) {
	db := s.DB()
	students := []model.Client{}

	return students, db.Where("class_id=? AND role=?", id, "student").Find(&students).Error
}

func (t *classRepo) DeleteByID(s repo.DBRepo, id string) (*model.Class, error) {
	db := s.DB()
	var param model.Class

	return &param, db.Where("id = ?", id).Delete(&param).Error
}
