package models

import "gorm.io/gorm"

type Paslon struct {
    gorm.Model
    Name 	string `json:"name" gorm:"type varchar(255)"`
    Visi 	string	`json:"visi" gorm:"type varchar(255)"`
    Image 	string `json:"image" gorm:"type varchar(255)"`
}


