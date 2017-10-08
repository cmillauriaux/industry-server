package model

type CompanyType int

const (
	FACTORY CompanyType = iota
	WAREHOUSE
	SHOP
	MINE
	FARM
	PLANT
)

type Company struct {
	CompanyID        string      `dynamo:"CompanyID"`
	PlayerID         string      `dynamo:"PlayerID"`
	Name             string      `dynamo:"Name"`
	CompanyType      CompanyType `dynamo:"CompanyType"`
	Country          string      `dynamo:"Country"`
	Latitude         int         `dynamo:"Latitude"`
	Longitude        int         `dynamo:"Longitude"`
	CurrentProductID string      `dynamo:"CurrentProductID"`
}
