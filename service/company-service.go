package service

import (
	"errors"

	"github.com/satori/go.uuid"

	"github.com/cmillauriaux/industry-server/model"
	"github.com/cmillauriaux/industry-server/persistance"
)

var db *persistance.CompanyPersistance

func init() {
	db = &persistance.CompanyPersistance{}
}

type CompanyService struct {
}

func (c *CompanyService) NewCompany(company *model.Company) error {
	// Check arguments
	if company == nil || company.PlayerID == uuid.Nil.String() || company.CompanyID == uuid.Nil.String() {
		return errors.New("Not enough arguments to create company")
	}

	// Check if player exists
	isPlayerExists, err := servicePlayer.IsPlayerExists(company.PlayerID)
	if err != nil {
		return err
	}
	if !isPlayerExists {
		return errors.New("Player doesn't exists")
	}

	// Create company
	err = db.New(company)
	return err
}

func (c *CompanyService) UpdateCompany(company *model.Company) error {
	isExists, err := c.IsCompanyExists(company.CompanyID)

	if err != nil {
		return err
	}

	if !isExists {
		return errors.New("Cannot find company")
	}

	err = db.Update(company)
	return err
}

func (c *CompanyService) UpdateProduction(playerID string, companyID string, productID string) error {
	// Check if player exists
	isPlayerExists, err := servicePlayer.IsPlayerExists(playerID)
	if err != nil {
		return err
	}
	if !isPlayerExists {
		return errors.New("Player doesn't exists")
	}

	// Check if company exists and player is owner
	isCompanyExists, err := c.IsPlayerCompany(playerID, companyID)
	if err != nil {
		return err
	}
	if !isCompanyExists {
		return errors.New("Company doesn't exists or player is not the owner")
	}

	// TODO: Check if the product exists

	// TODO: Check if the product can be made in the company

	// Get the company
	company, err := c.GetCompany(companyID)
	if err != nil {
		return err
	}

	// Set the new product
	company.CurrentProductID = productID

	// Create company
	err = db.Update(company)
	return err
}

func (c *CompanyService) GetCompany(companyID string) (*model.Company, error) {
	company, err := db.Get(companyID)
	if company == nil {
		return nil, errors.New("Cannot find this company")
	}
	return company, err
}

func (c *CompanyService) GetCompanyByPlayerID(playerID string) ([]*model.Company, error) {
	return db.GetByPlayerID(playerID)
}

func (c *CompanyService) IsPlayerCompany(playerID string, companyID string) (bool, error) {
	company, err := c.GetCompany(companyID)
	if err != nil {
		return false, err
	}
	if company == nil || company.CompanyID == "" {
		return false, nil
	}
	if company.PlayerID == playerID {
		return true, nil
	}
	return false, nil
}

func (c *CompanyService) IsCompanyExists(companyID string) (bool, error) {
	company, err := db.Get(companyID)
	if company == nil {
		return false, errors.New("Cannot find this company")
	}
	if company != nil && company.CompanyID != uuid.Nil.String() {
		return true, nil
	}
	return false, err
}
