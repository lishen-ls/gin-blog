package models

import "time"

type Id struct {
	Id uint `json:"id" gorm:"primaryKey"`
}

type TimeStamp struct {
	Createtime time.Time `json:"createtime"`
	Updatetime time.Time `json:"updatetime"`
}
