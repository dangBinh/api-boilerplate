package storage

// Customer ...
type Customer struct {
	ID         int `gorm:"primary_key"`
	Name       string
	Tel        string
	PassportNo string `json:"passport" gorm:"column:passport_no"`
}

// CustomerStorage ...
type CustomerStorage interface {
	storageCRUD
}

type customerstorageImpl struct {
	*storageImpl
}

// TableName ...
func (m Customer) TableName() string {
	return "customer"
}

// NewCustomerStorage ...
func NewCustomerStorage() CustomerStorage {
	return &customerstorageImpl{pkgDAO}
}
