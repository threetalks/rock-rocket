package models

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name        string `gorm:"index:idx_name"`
	Code        string `gorm:"code uniqueIndex:uidx_code"`
	LegalPerson string `gorm:"index:idx_legal_person"`
	Province    string `gorm:"index:idx_province"`
	City        string `gorm:"index:idx_city"`
	County      string `gorm:"index:idx_county"`
	Registered  string `gorm:"index:idx_registered"`
	Capital     int    `gorm:"index:idx_capital,default:-1"`
	Contact     string `gorm:"index:idx_contact,default:''"`
	Mail        string `gorm:"mail,default:''"`
	URL         string `gorm:"url,default:''"`
	Scope       string `gorm:"scope"`
}

func (company *Company) TableName() string {
	return "companies"
}
