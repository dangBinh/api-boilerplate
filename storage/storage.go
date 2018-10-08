package storage

import (
	"api-boilerplate/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql
)

type storageCRUD interface {
	ByID(int) (interface{}, error)
	List()
	Create(interface{}) error
	Update(interface{}) error
	Delete()
}

type storageImpl struct {
	db *gorm.DB
}

var pkgDAO *storageImpl

// GetStorage object
func init() {
	conf := config.GetConfig()
	db, err := gorm.Open("mysql", conf.MySQLConnectionString)
	if err != nil {
		panic(err)
	}
	pkgDAO = &storageImpl{db}
}

func (s *storageImpl) ByID(id int) (obj interface{}, err error) {
	dbc := s.db.First(&obj, id)
	if dbc.Error != nil {
		return nil, dbc.Error
	}
	return obj, nil
}

func (s *storageImpl) List() {

}

func (s *storageImpl) Create(data interface{}) error {
	dbc := s.db.Create(data)
	if dbc.Error != nil {
		return dbc.Error
	}
	return nil
}

func (s *storageImpl) Update(data interface{}) error {
	dbc := s.db.Save(data)
	if dbc.Error != nil {
		return dbc.Error
	}
	return nil
}

func (s *storageImpl) Delete() {

}
