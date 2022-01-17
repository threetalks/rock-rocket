package rock_rocket

import (
	"reflect"
	"rock-rocket/internal/models"
	"testing"
)

func TestListCompanies(t *testing.T) {
	type args struct {
		legal      *string
		contact    *string
		province   *string
		city       *string
		county     *string
		capital    *string
		registered *string
		industry   *string
		typ        *string
		address    *string
		scope      *string
		offset     int
		limit      int
	}
	tests := []struct {
		name    string
		args    args
		want    []models.Company
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListCompanies(tt.args.legal, tt.args.contact, tt.args.province, tt.args.city, tt.args.county, tt.args.capital, tt.args.registered, tt.args.industry, tt.args.typ, tt.args.address, tt.args.scope, tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListCompanies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListCompanies() got = %v, want %v", got, tt.want)
			}
		})
	}
}
