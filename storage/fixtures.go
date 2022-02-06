package storage

import "github.com/go-rest-playground/models"

var classes = []*models.Class{
	{ID: "PI001", Name: "Pilates", StartDate: createTime("2020-01-29"), EndDate: createTime("2020-02-28"), Capacity: 20},
	{ID: "D+001", Name: "Dance+", StartDate: createTime("2020-01-29"), EndDate: createTime("2020-02-28"), Capacity: 20},
	{ID: "FB001", Name: "Full Body", StartDate: createTime("2020-01-29"), EndDate: createTime("2020-02-28"), Capacity: 20},
	{ID: "YO001", Name: "Yoga", StartDate: createTime("2020-01-29"), EndDate: createTime("2020-02-28"), Capacity: 20},
}
