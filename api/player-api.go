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

func NewPlayer(c *gin.Context) {
	// Validate datas
	var player model.Player
	err := c.BindJSON(&player)
	if err != nil {
		c.JSON(http.StatusBadRequest, structToJSON(makeError(err)))
		return
	}

	// Make a new UID
	player.PlayerID = uuid.NewV4().String()

	// Register player
	err = playerService.NewPlayer(&player)
	if err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, structToJSON(makeError(err)))
			return
		}
	}

	makeResponse(c, &player, &output.PlayerOutput{})
}

func UpdatePlayer(c *gin.Context) {
	// Validate datas
	var player model.Player
	err := c.BindJSON(&player)
	if err != nil {
		c.JSON(http.StatusBadRequest, structToJSON(makeError(err)))
		return
	}

	// Update player
	err = playerService.UpdatePlayer(&player)
	if err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, structToJSON(makeError(err)))
			return
		}
	}

	makeResponse(c, &player, &output.PlayerOutput{})
}

func GetPlayer(c *gin.Context) {
	// get datas
	playerID := c.Param("id")

	// Register player
	player, err := playerService.GetPlayer(playerID)
	if err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, structToJSON(makeError(err)))
			return
		}
	}

	makeResponse(c, player, &output.PlayerOutput{})
}
