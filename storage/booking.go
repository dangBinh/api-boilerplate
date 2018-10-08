package storage

//go:generate mockery -case underscore -name BookingStorage -inpkg

// Booking ...
type Booking struct {
	ID       int    `gorm:"primary_key"`
	Cid      int    `gorm:"column:c_id"`
	Checkin  string `json:"checkin"`
	Checkout string `json:"checkout"`
	Status   int    `json:"status"`
	Paid     int    `json:"paid"`
}

// BookingStorage ...
type BookingStorage interface {
	storageCRUD
}

type bookingStorageImpl struct {
	*storageImpl
}

// TableName ...
func (m Booking) TableName() string {
	return "booking"
}

// NewBookingStorage ...
func NewBookingStorage() BookingStorage {
	return &bookingStorageImpl{pkgDAO}
}
