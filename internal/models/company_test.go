/*
Copyright © 2022 Three Talks  <bytebody@icloud.com>
*/

package models

import (
	"testing"
)

func TestCompany_TableName(t *testing.T) {
	company := &Company{}
	want := "companies"
	if got := company.TableName(); got != want {
		t.Errorf("TableName() = %v, want %v", got, want)
	}
}
