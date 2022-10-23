package entities

import (
	"sync"
)

type CovidUsecase interface {
	Summary(req *CovidDataReq) (*CovidSummaryRes, error)
}

type CovidRepository interface {
	GetCovidData(req *CovidDataReq) (*CovidDataRes, error)
}

type CovidDataReq struct{}

type CovidData struct {
	ConfirmDate    string
	No             int
	Age            int
	Gender         string
	GenderEn       string
	Nation         string
	NationEn       string
	Province       string
	ProvinceId     int
	District       string
	ProvinceEn     string
	StatQuarantine int
}

type CovidDataRes struct {
	Data []CovidData `json:"data"`
}

type CovidSummaryRes struct {
	Province map[string]int `json:"province"`
	AgeGroup map[string]int `json:"ageGroup"`
	Mu       sync.Mutex
}

func (c *CovidSummaryRes) CountProvince(province string) {
	c.Mu.Lock()
	c.Province[province] += 1
	c.Mu.Unlock()
}

func (c *CovidSummaryRes) CountAgeGroup(ageGroup string) {
	c.Mu.Lock()
	c.AgeGroup[ageGroup] += 1
	c.Mu.Unlock()
}

func (c *CovidSummaryRes) CountData(val CovidData) {
	// prevent race condition
	c.Mu.Lock()
	if val.ProvinceEn != "" {
		c.Province[val.ProvinceEn] += 1
	} else {
		c.Province["Unknown Province"] += 1
	}

	if val.Age == 0 {
		c.AgeGroup["N/A"] += 1
	} else if val.Age >= 1 && val.Age <= 30 {
		c.AgeGroup["0-30"] += 1
	} else if val.Age > 30 && val.Age <= 60 {
		c.AgeGroup["31-60"] += 1
	} else if val.Age > 60 {
		c.AgeGroup["61+"] += 1
	}
	c.Mu.Unlock()
}
