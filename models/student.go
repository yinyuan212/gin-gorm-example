package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name   string
	Gender string
	Age    uint
}
