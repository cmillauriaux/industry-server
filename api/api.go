package api

import (
	"git.icysoft.fr/cedric/industry-go-server/service"
)

var playerService *service.PlayerService
var companyService *service.CompanyService

func init() {
	playerService = &service.PlayerService{}
	companyService = &service.CompanyService{}
}
