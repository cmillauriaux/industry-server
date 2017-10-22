package persistance

import (
	"git.icysoft.fr/cedric/industry-go-server/model"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

const PlayerTableName = "player"

var dbPlayer *dynamo.DB

func init() {
	dbPlayer = dynamo.New(session.New())
}

type PlayerPersistance struct {
}

func (c *PlayerPersistance) New(player *model.Player) error {
	table := dbPlayer.Table(PlayerTableName)
	return table.Put(player).Run()
}

func (c *PlayerPersistance) Update(player *model.Player) error {
	table := dbPlayer.Table(PlayerTableName)
	return table.Put(player).Run()
}

func (c *PlayerPersistance) Delete(playerID string) error {
	table := dbPlayer.Table(PlayerTableName)
	return table.Delete("PlayerID", playerID).Run()
}

func (c *PlayerPersistance) Get(playerID string) (*model.Player, error) {
	var result model.Player
	table := dbPlayer.Table(PlayerTableName)
	err := table.Get("PlayerID", playerID).One(&result)
	return &result, err
}

func (c *PlayerPersistance) GetByMail(mail string) (*model.Player, error) {
	var result model.Player
	table := dbPlayer.Table(PlayerTableName)
	err := table.Get("Mail", mail).One(&result)
	return &result, err
}

func (c *PlayerPersistance) List() ([]*model.Player, error) {
	var results []*model.Player
	table := dbPlayer.Table(PlayerTableName)
	err := table.Scan().All(&results)
	return results, err
}
