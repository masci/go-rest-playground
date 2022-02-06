package storage

import "github.com/masci/go-rest-playground/models"

var classes = []*models.Class{
	{ID: "PI0001", Name: "Pilates", StartDate: createTime("2020-01-29"), EndDate: createTime("2020-02-28"), Capacity: 20},
	{ID: "DA0001", Name: "Dance+", StartDate: createTime("2020-01-29"), EndDate: createTime("2020-02-28"), Capacity: 20},
	{ID: "FB0001", Name: "Full Body", StartDate: createTime("2020-01-29"), EndDate: createTime("2020-02-28"), Capacity: 20},
	{ID: "YO0001", Name: "Yoga", StartDate: createTime("2020-01-29"), EndDate: createTime("2020-02-28"), Capacity: 20},
}
