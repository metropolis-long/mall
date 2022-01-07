package data

import (
	"fmt"
	"strings"

	"time"

	"github.com/spaolacci/murmur3"
)

func (l *Like) Add(ip string, ua string, title string) error {
	db := Db
	if Db == nil {
		panic("databases init fail")
	}
	like := &Like{
		Ip:        ip,
		Ua:        ua,
		Title:     title,
		Hash:      murmur3.Sum64([]byte(strings.Join([]string{ip, ua, title}, "-"))) >> 1,
		CreatedAt: time.Now(),
	}
	if err := db.Create(like).Error; err != nil {
		return err
	}
	return nil
}

func (l *Like) SumLikeByPerson() {
	db := Db
	results := []Result{}
	err := db.Table("cities_city").Select("province_code , count(province_code) as total").Group("province_code").Having("count(province_code) > ?", 10).Scan(&results).Error
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("2222222222222")
	for _,r:=range results{
		fmt.Println(r)
	}

	// fmt.Fprintln()
	// rows, err := db.Table("cities_city").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Rows()
	// 	for rows.Next() {
	// rows.Columns()
	// }
	// results := []Result{}
	// db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Scan(&results)
}
