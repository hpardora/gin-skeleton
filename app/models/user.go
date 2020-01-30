package models

import "time"

type User struct {
	Id        int       `gorm:"primary_key" json:"id"`
	Account   string    `gorm:"type:varchar(50) not null comment '账号'"`
	Password  string    `gorm:"type:varchar(255) not null comment '密码'"`
	CreatedAt time.Time `gorm:"type:timestamp not null" json:"created_at"`
}
