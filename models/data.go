package models

import "github.com/jinzhu/gorm"

type Data struct {
	gorm.Model

	Texto string `gorm:"column:texto"`
	// ... Otras columnas de la tabla
}


