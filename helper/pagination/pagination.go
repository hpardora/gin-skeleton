package pagination

import (
	"github.com/jinzhu/gorm"
)

type Param struct {
	DB    *gorm.DB
	Page  int
	Limit int
}

type Paginator struct {
	Data
	Page
}

type Page struct {
	Total   int `json:"total"`
	PerPage int `json:"per_page"`
}

type Data struct {
	Data interface{} `json:"data"`
}

func Paging(p *Param, result interface{}) *Paginator {
	db := p.DB

	if p.Page < 1 {
		p.Page = 1
	}

	if p.Limit == 0 {
		p.Limit = 15
	}

	done := make(chan bool, 1)
	var paginator Paginator
	var count int
	var offset int

	go countRecords(db, result, done, &count)

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.Limit
	}

	db.Limit(p.Limit).Offset(offset).Find(result)
	<-done

	paginator.Data.Data = result
	paginator.Page.Total = count
	paginator.Page.PerPage = p.Limit

	return &paginator
}

func countRecords(db *gorm.DB, anyType interface{}, done chan bool, count *int) {
	db.Model(anyType).Count(count)
	done <- true
}
