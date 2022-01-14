package company

import (
	"reflect"
	"rock-rocket/internal/models"
	"testing"
)

func TestGetCompany(t *testing.T) {
	type args struct {
		id   *string
		name *string
		code *string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Company
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCompany(tt.args.id, tt.args.name, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCompany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCompany() got = %v, want %v", got, tt.want)
			}
		})
	}
}
