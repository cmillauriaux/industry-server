package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"

	"github.com/cmillauriaux/indeustry-server/model"
	"github.com/cmillauriaux/indeustry-server/model/output"
)

func init() {

}

func NewCompany(c *gin.Context) {
	// Validate datas
	var company model.Company
	err := c.BindJSON(&company)
	if err != nil {
		c.JSON(http.StatusBadRequest, structToJSON(makeError(err)))
		return
	}

	// Make a new UID
	company.CompanyID = uuid.NewV4().String()

	// Register company
	err = companyService.NewCompany(&company)
	if err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, structToJSON(makeError(err)))
			return
		}
	}

	makeResponse(c, &company, &output.CompanyOutput{})
}

func UpdateCompany(c *gin.Context) {
	// Validate datas
	var company model.Company
	err := c.BindJSON(&company)
	if err != nil {
		c.JSON(http.StatusBadRequest, structToJSON(makeError(err)))
		return
	}

	// Update company
	err = companyService.UpdateCompany(&company)
	if err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, structToJSON(makeError(err)))
			return
		}
	}

	makeResponse(c, &company, &output.CompanyOutput{})
}

func GetCompany(c *gin.Context) {
	// get datas
	companyID := c.Param("id")

	// Register player
	company, err := companyService.GetCompany(companyID)
	if err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, structToJSON(makeError(err)))
			return
		}
	}

	makeResponse(c, &company, &output.CompanyOutput{})
}
