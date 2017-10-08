package model

type ProductType int

const (
	TOOL ProductType = iota
	RAWMATERIAL
	PRODUCT
	GOOD
	ENERGY
)

type Product struct {
	ProductID         string      `dynamo:"productID"`
	Name              string      `dynamo:"name"`
	ProductType       ProductType `dynamo:"productType"`
	CanBeMadeIn       CompanyType `dynamo:"canBeMadeIn"`
	RawMaterials      []string    `dynamo:",set"`
	AverageTimeToMake int         `dynamo:"averageTimeToMake"`
}
