package persistance

import (
	"reflect"
	"testing"

	"github.com/cmillauriaux/industry-server/model"
	uuid "github.com/satori/go.uuid"
)

func TestCompanyPersistance_New(t *testing.T) {
	type args struct {
		company *model.Company
	}
	tests := []struct {
		name    string
		c       *CompanyPersistance
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"insertOK", &CompanyPersistance{}, args{company: &model.Company{CompanyID: uuid.NewV4().String(), PlayerID: uuid.NewV4().String(), Country: "TEST", Name: "TEST", Latitude: 45, Longitude: 15}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CompanyPersistance{}
			if err := c.New(tt.args.company); (err != nil) != tt.wantErr {
				t.Errorf("CompanyPersistance.New() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCompanyPersistance_Update(t *testing.T) {
	type args struct {
		company *model.Company
	}
	tests := []struct {
		name    string
		c       *CompanyPersistance
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CompanyPersistance{}
			if err := c.Update(tt.args.company); (err != nil) != tt.wantErr {
				t.Errorf("CompanyPersistance.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCompanyPersistance_Delete(t *testing.T) {
	type args struct {
		companyID string
	}
	tests := []struct {
		name    string
		c       *CompanyPersistance
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CompanyPersistance{}
			if err := c.Delete(tt.args.companyID); (err != nil) != tt.wantErr {
				t.Errorf("CompanyPersistance.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCompanyPersistance_Get(t *testing.T) {
	type args struct {
		companyID string
	}
	tests := []struct {
		name    string
		c       *CompanyPersistance
		args    args
		want    *model.Company
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CompanyPersistance{}
			got, err := c.Get(tt.args.companyID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompanyPersistance.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompanyPersistance.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompanyPersistance_GetByPlayerID(t *testing.T) {
	type args struct {
		playerID string
	}
	tests := []struct {
		name    string
		c       *CompanyPersistance
		args    args
		want    []*model.Company
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CompanyPersistance{}
			got, err := c.GetByPlayerID(tt.args.playerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompanyPersistance.GetByPlayerID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompanyPersistance.GetByPlayerID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompanyPersistance_List(t *testing.T) {
	tests := []struct {
		name    string
		c       *CompanyPersistance
		want    []*model.Company
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CompanyPersistance{}
			got, err := c.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("CompanyPersistance.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompanyPersistance.List() = %v, want %v", got, tt.want)
			}
		})
	}
}
