package persistance

import (
	"github.com/cmillauriaux/indeustry-server/model"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

const PlayerTableName = "player"

var dbPlayer *dynamo.DB
var tablePlayer *dynamo.Table

func init() {
	dbPlayer = dynamo.New(session.New())
	tablePlayerNoPtr := dbPlayer.Table(PlayerTableName)
	tablePlayer = &tablePlayerNoPtr
}

type PlayerPersistance struct {
}

func (c *PlayerPersistance) New(player *model.Player) error {
	return tablePlayer.Put(player).Run()
}

func (c *PlayerPersistance) Update(player *model.Player) error {
	return tablePlayer.Put(player).Run()
}

func (c *PlayerPersistance) Delete(playerID string) error {
	return tablePlayer.Delete("PlayerID", playerID).Run()
}

func (c *PlayerPersistance) Get(playerID string) (*model.Player, error) {
	var result model.Player
	err := tablePlayer.Get("PlayerID", playerID).One(&result)
	return &result, err
}

func (c *PlayerPersistance) GetByMail(mail string) (*model.Player, error) {
	var result model.Player
	err := tablePlayer.Get("Mail", mail).One(&result)
	return &result, err
}

func (c *PlayerPersistance) List() ([]*model.Player, error) {
	var results []*model.Player
	err := tablePlayer.Scan().All(&results)
	return results, err
}
