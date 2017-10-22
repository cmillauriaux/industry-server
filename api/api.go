package api

import (
	"git.icysoft.fr/cedric/industry-go-server/service"
)

var playerService *service.PlayerService

func init() {
	playerService = &service.PlayerService{}
}
