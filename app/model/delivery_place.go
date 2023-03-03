package model

import (
	"time"
)

type Example struct {
	Id             int64      `json:"place_id"`
	AppId          string     `json:"app_id,omitempty"`
	PlaceName      string     `json:"place_name"`
	PlacePhone     string     `json:"place_phone"`
	Province       string     `json:"province"`
	City           string     `json:"city"`
	County         string     `json:"county"`
	Detail         string     `json:"detail"`
	Address        string     `json:"address"`
	OperatingCycle string     `json:"operating_cycle"`
	BusinessHours  string     `json:"business_hours"`
	IsDefault      int8       `json:"is_default"`
	Remark         string     `json:"-"`
	IsDeleted      int        `json:"-"`
	CreatedAt      *time.Time `json:"-"`
	UpdatedAt      *time.Time `json:"-"`
}

func (t *Example) TableName() string {
	return "t_example"
}

func (t *Example) DeleteExample() int64 {
	result := ExampleDB.
		Table(t.TableName()).
		Where("id = ?", t.Id).
		Update("is_deleted", 0)
	return result.RowsAffected
}

func (t *Example) GetExampleDetail(placeId int64) (e Example, err error) {
	err = ExampleDB.
		Select("id,app_id,place_name,place_phone,province,city,county,detail,address,is_default,operating_cycle,business_hours").
		Where("id = ?", placeId).
		Where("is_deleted = ?", 0).
		First(&e).Error
	if err != nil {
		return e, err
	}

	return e, nil
}

func (t *Example) CountExample() (count int64, err error) {
	err = ExampleDB.
		Table(t.TableName()).
		Where("app_id = ?", t.AppId).
		Where("is_deleted = ?", 0).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetList ,,,
func (t *Example) GetList(ids []string) (list []Example, err error) {
	err = ExampleDB.
		Select("id,address").
		Where("id in (?)", ids).
		Where("is_deleted = ?", 0).
		Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}
