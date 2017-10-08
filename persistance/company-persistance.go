package persistance

import (
	"git.icysoft.fr/cedric/industry-go-server/model"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

const CompanyTableName = "company"

var db *dynamo.DB

func init() {
	db = dynamo.New(session.New())
}

type CompanyPersistance struct {
}

func (c *CompanyPersistance) New(company *model.Company) error {
	table := db.Table(CompanyTableName)
	return table.Put(company).Run()
}

func (c *CompanyPersistance) Update(company *model.Company) error {
	table := db.Table(CompanyTableName)
	return table.Update(company.CompanyID, company).Run()
}

func (c *CompanyPersistance) Delete(companyID string) error {
	table := db.Table(CompanyTableName)
	return table.Delete("CompanyID", companyID).Run()
}

func (c *CompanyPersistance) Get(companyID string) (*model.Company, error) {
	var result *model.Company
	table := db.Table(CompanyTableName)
	err := table.Get("CompanyID", companyID).One(result)
	return result, err
}

func (c *CompanyPersistance) GetByPlayerID(playerID string) ([]*model.Company, error) {
	var results []*model.Company
	table := db.Table(CompanyTableName)
	err := table.Get("PlayerID", playerID).All(results)
	return results, err
}

func (c *CompanyPersistance) List() ([]*model.Company, error) {
	var results []*model.Company
	table := db.Table(CompanyTableName)
	err := table.Scan().All(&results)
	return results, err
}
