package storage

import (
	"api-boilerplate/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql
)

type storageCRUD interface {
	ByID()
	List()
	Create(interface{})
	Update()
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

func (s *storageImpl) ByID() {

}

func (s *storageImpl) List() {

}

func (s *storageImpl) Create(obj interface{}) {
	s.db.Create(obj)
}

func (s *storageImpl) Update() {

}

func (s *storageImpl) Delete() {

}
