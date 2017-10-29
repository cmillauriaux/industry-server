package api

import (
	"github.com/cmillauriaux/industry-server/service"
)

var playerService *service.PlayerService
var companyService *service.CompanyService

func init() {
	playerService = &service.PlayerService{}
	companyService = &service.CompanyService{}
}
