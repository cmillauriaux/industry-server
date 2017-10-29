package persistance

import (
	"github.com/cmillauriaux/indeustry-server/model"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

const CompanyTableName = "company"

var db *dynamo.DB
var tableCompany *dynamo.Table

func init() {
	db = dynamo.New(session.New())
	tableCompanyNoPtr := dbPlayer.Table(PlayerTableName)
	tableCompany = &tableCompanyNoPtr
}

type CompanyPersistance struct {
}

func (c *CompanyPersistance) New(company *model.Company) error {
	return tableCompany.Put(company).Run()
}

func (c *CompanyPersistance) Update(company *model.Company) error {
	return tableCompany.Update(company.CompanyID, company).Run()
}

func (c *CompanyPersistance) Delete(companyID string) error {
	return tableCompany.Delete("CompanyID", companyID).Run()
}

func (c *CompanyPersistance) Get(companyID string) (*model.Company, error) {
	var result *model.Company
	err := tableCompany.Get("CompanyID", companyID).One(result)
	return result, err
}

func (c *CompanyPersistance) GetByPlayerID(playerID string) ([]*model.Company, error) {
	var results []*model.Company
	err := tableCompany.Get("PlayerID", playerID).All(results)
	return results, err
}

func (c *CompanyPersistance) List() ([]*model.Company, error) {
	var results []*model.Company
	err := tableCompany.Scan().All(&results)
	return results, err
}
