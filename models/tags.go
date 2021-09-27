package models

import "time"

type Tag struct {
	ID uint `json:"id" gorm:"primary_key"`
	Name  string  `json:"name"`
}

type TagRelation struct {
	ID uint `json:"id" gorm:"primary_key"`
	UserID uint
	Tags []Tag `gorm:"many2many:tag_relation_tags;"`
	Expire time.Time
}