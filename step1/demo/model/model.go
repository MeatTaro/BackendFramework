package model

import (
	"database/sql"
	// "time"
)

type Process struct {
	Id string `json:"id"`

	ProductType   string `json:"producttype"`
	Characteristic string `json:"characteristic"`

	// CreateTime time.Time `json:"createTime"`
	// UpdateTime time.Time `json:"updateTime"`
}

var (
	//宣告變數作為sql接口
	//type DB: https://golang.org/pkg/database/sql/#DB
	Db *sql.DB
)

func GetAll() ([]Process, error)  {
	//create the object slice
	result := []Process{}

	//load the object data from the database
	rows, err := Db.Query("SELECT id, productType, characteristic FROM cats order by id desc")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item Process
		if err := rows.Scan(&item.Id, &item.ProductType, &item.Characteristic); err != nil {
			return nil, err
		}
		result = append(result, item)
	}

	return result, err
}
