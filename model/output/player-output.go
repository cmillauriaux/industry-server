package output

type PlayerOutput struct {
	PlayerID    string `dynamo:"PlayerID" json:"player_id,omitempty"`
	Mail        string `dynamo:"Mail" json:"mail,omitempty"`
	Name        string `dynamo:"Name" json:"name,omitempty"`
	Nationality string `dynamo:"Nationality" json:"nationality,omitempty"`
}
