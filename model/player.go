package model

type Player struct {
	PlayerID    string `dynamo:"PlayerID" json:"player_id,omitempty"`
	Mail        string `dynamo:"Mail" json:"mail,omitempty"`
	Password    string `dynamo:"Password" json:"password,omitempty"`
	Name        string `dynamo:"Name" json:"name,omitempty"`
	Nationality string `dynamo:"Nationality" json:"nationality,omitempty"`
}
