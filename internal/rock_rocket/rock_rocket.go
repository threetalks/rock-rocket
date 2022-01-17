/*
Copyright © 2022 Three Talks  <bytebody@icloud.com>
*/

package rock_rocket

import (
	"fmt"
	"gorm.io/gorm"
	"rock-rocket/internal/db"
	"rock-rocket/internal/models"
)

type RockRocket struct {
	DB *gorm.DB
}

func (rocket *RockRocket) GetCompanyByID(id uint) (*models.Company, error) {
	var company *models.Company
	rocket.DB.First(company, "id=?", id)
	if company == nil {
		return nil, fmt.Errorf("not found")
	}
	return company, nil
}

func (rocket *RockRocket) GetCompanyByCode(code string) (*models.Company, error) {
	var company *models.Company
	db.DB().First(company, "code=?", code)
	if company == nil {
		return nil, fmt.Errorf("not found")
	}
	return company, nil
}

func (rocket *RockRocket) GetCompanyByName(name string) (*models.Company, error) {
	var company *models.Company
	rocket.DB.First(company, "name=?", name)
	if company == nil {
		return nil, fmt.Errorf("not found")
	}
	return company, nil
}

func (rocket *RockRocket) ListCompanies() ([]models.Company, error) {
	return nil, nil
}

func NewRockRocket(db *gorm.DB) *RockRocket {
	return &RockRocket{
		DB: db,
	}
}
