package persistance

import (
	"git.icysoft.fr/cedric/industry-go-server/model"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

const productTableName = "product"

var dbProduct *dynamo.DB

func init() {
	dbProduct = dynamo.New(session.New())
}

type ProductPersistance struct {
}

func (c *ProductPersistance) New(product *model.Product) error {
	table := dbProduct.Table(productTableName)
	return table.Put(product).Run()
}

func (c *ProductPersistance) Update(product *model.Product) error {
	table := dbProduct.Table(productTableName)
	return table.Update(product.ProductID, product).Run()
}

func (c *ProductPersistance) Delete(productID string) error {
	table := dbProduct.Table(productTableName)
	return table.Delete("ProductID", productID).Run()
}

func (c *ProductPersistance) Get(productID string) (*model.Product, error) {
	var result *model.Product
	table := dbProduct.Table(productTableName)
	err := table.Get("ProductID", productID).One(result)
	return result, err
}

func (c *ProductPersistance) List() ([]*model.Product, error) {
	var results []*model.Product
	table := dbProduct.Table(productTableName)
	err := table.Scan().All(&results)
	return results, err
}
