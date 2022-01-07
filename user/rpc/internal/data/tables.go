package data

import "time"

type Like struct {
	ID        int    `gorm:"primary_key"`
	UserId    int    `gorm:"type:int"`
	Ip        string `gorm:"type:varchar(20);not null;index:ip_idx"`
	Ua        string `gorm:"type:varchar(256);not null;"`
	Title     string `gorm:"type:varchar(128);not null;index:title_idx"`
	Hash      uint64 `gorm:"unique_index:hash_idx;"`
	CreatedAt time.Time
}

type Users struct {
	UserId    int    `gorm:"primary_key"`
	Name      string `gorm:"type:varchar(20);not null;`
	Phone     string `gorm:"type:varchar(20);not null;"`
	Group     uint64 `gorm:"type:int;"`
	CreatedAt time.Time
}
type Result struct {
	ProvinceCode int64
	Total        int64
}
type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}
