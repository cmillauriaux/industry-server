package model

type Storage struct {
	StorageID       string
	CompanyID       string
	ProductID       string
	ProductType     ProductType
	Quantity        int
	Price           int
	Quality         int
	MaintenanceCost int
}
