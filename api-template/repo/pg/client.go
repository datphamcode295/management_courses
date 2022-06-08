package pg

import (
	"github.com/MStation-io/api/model"
	"github.com/MStation-io/api/repo"
)

type clientRepo struct{}

func NewClientRepo() repo.ClientRepo {
	return &clientRepo{}
}

func (t *clientRepo) GetAllClients(s repo.DBRepo) ([]model.Client, error) { //

	clients := []model.Client{}
	db := s.DB()

	return clients, db.Find(&clients).Error
}

func (t *clientRepo) PostClient(s repo.DBRepo, param model.Client) (*model.Client, error) {

	db := s.DB()

	x := db.Create(&param)
	return &param, x.Error
}

func (t *clientRepo) Delete(s repo.DBRepo, id string) (*model.Client, error) {
	db := s.DB()
	var temp model.Client

	return &temp, db.Where("id = ?", id).Delete(&temp).Error
}

func (t *clientRepo) UpdatePassword(s repo.DBRepo, id string, password string) (*model.Client, error) {
	db := s.DB()
	client := model.Client{}
	return &client, db.Model(&client).Where("id = ?", id).Update("password", password).Error
}

func (t *clientRepo) CheckingLogin(s repo.DBRepo, username string, password string) (*model.Client, error) {
	db := s.DB()
	client := model.Client{}

	return &client, db.Where("username = ? AND password = ?", username, password).First(&client).Error
}

// func (t *classRepo) GetClassByID(s repo.DBRepo, id string) (*model.Class, error) {
// 	db := s.DB()
// 	var class model.Class

// 	return &class, db.Find(&class, "id=?", id).Error
// }

// func (t *classRepo) GetClassByName(s repo.DBRepo, name string) (*model.Class, error) {
// 	db := s.DB()
// 	var class model.Class

// 	return &class, db.Find(&class, "name=?", name).Error
// }
