package output

type CompanyOutput struct {
	CompanyID   string `dynamo:"CompanyID" json:"company_id,omitempty"`
	Mail        string `dynamo:"Mail" json:"mail,omitempty"`
	Name        string `dynamo:"Name" json:"name,omitempty"`
	Nationality string `dynamo:"Nationality" json:"nationality,omitempty"`
}
