package service

import (
	"errors"

	"github.com/satori/go.uuid"

	"git.icysoft.fr/cedric/industry-go-server/model"
	"git.icysoft.fr/cedric/industry-go-server/persistance"
)

var dbPlayer *persistance.PlayerPersistance
var servicePlayer *PlayerService

func init() {
	dbPlayer = &persistance.PlayerPersistance{}
	servicePlayer = &PlayerService{}
}

type PlayerService struct {
}

func (c *PlayerService) NewPlayer(player *model.Player) error {
	err := dbPlayer.New(player)
	return err
}

func (c *PlayerService) UpdatePlayer(player *model.Player) error {
	isExists, err := c.IsPlayerExists(player.PlayerID)

	if err != nil {
		return err
	}

	if !isExists {
		return errors.New("Cannot find player")
	}

	err = dbPlayer.Update(player)
	return err
}

func (c *PlayerService) GetPlayer(playerID string) (*model.Player, error) {
	player, err := dbPlayer.Get(playerID)
	if err != nil {
		return nil, err
	}
	if player == nil {
		return nil, errors.New("Cannot find this player")
	}
	return player, err
}

func (c *PlayerService) IsPlayerExists(playerID string) (bool, error) {
	player, err := dbPlayer.Get(playerID)
	if player == nil {
		return false, errors.New("Cannot find this player")
	}
	if player != nil && player.PlayerID != uuid.Nil.String() {
		return true, nil
	}
	return false, err
}
