package usecases

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/nuttapon-first/summarize-covid19-stats/modules/entities"
)

// type fakeCovidiDataRes struct {
// 	Data string
// }

// type CovidRepository interface {
// 	GetCovidData(req *entities.CovidDataReq) (*fakeCovidiDataRes, error)
// }

type fakeCovidRepositorySuccess struct{}

func (f *fakeCovidRepositorySuccess) GetCovidData(req *entities.CovidDataReq) (*entities.CovidDataRes, error) {
	covidData := &entities.CovidDataRes{}
	data := `{"Data":[{"ConfirmDate":"2021-05-04","No":null,"Age":51,"Gender":"หญิง","GenderEn":"Female","Nation":null,"NationEn":"China","Province":"Phrae","ProvinceId":46,"District":null,"ProvinceEn":"Phrae","StatQuarantine":5},{"ConfirmDate":"2021-05-01","No":null,"Age":51,"Gender":"ชาย","GenderEn":"Male","Nation":null,"NationEn":"India","Province":"Suphan Buri","ProvinceId":65,"District":null,"ProvinceEn":"Suphan Buri","StatQuarantine":8},{"ConfirmDate":"2021-05-01","No":null,"Age":79,"Gender":null,"GenderEn":null,"Nation":null,"NationEn":"India","Province":"Roi Et","ProvinceId":53,"District":null,"ProvinceEn":"Roi Et","StatQuarantine":1}]}`
	json.Unmarshal([]byte(data), covidData)
	return covidData, nil
}

func TestSummaryCovidDataSuccess(t *testing.T) {
	fakeCovidRepository := &fakeCovidRepositorySuccess{}
	covidUses := NewCovidUsecase(fakeCovidRepository)
	req := &entities.CovidDataReq{}
	fakeData, err := covidUses.Summary(req)

	if err != nil {
		t.Errorf("expected success but got error %q\n", err)
	}

	want := &entities.CovidSummaryRes{
		Province: map[string]int{"Phrae": 1, "Roi Et": 1, "Suphan Buri": 1},
		AgeGroup: map[string]int{"0-30": 0, "31-60": 2, "61+": 1, "N/A": 0},
	}

	provinceSum := 0
	for province := range fakeData.Province {
		provinceSum += fakeData.Province[province]
		if fakeData.Province[province] != want.Province[province] {
			t.Errorf("%s : %d is expected but got %d\n", province, want.Province[province], fakeData.Province[province])
		}
	}

	ageGroupSum := 0
	for ageGroup := range fakeData.AgeGroup {
		ageGroupSum += fakeData.AgeGroup[ageGroup]
		if fakeData.AgeGroup[ageGroup] != want.AgeGroup[ageGroup] {
			t.Errorf("%s : %d is expected but got %d\n", ageGroup, want.AgeGroup[ageGroup], fakeData.AgeGroup[ageGroup])
		}
	}

	if provinceSum != ageGroupSum {
		t.Errorf("province and age group should equal each other [province: %d] [age group: %d]\n", provinceSum, ageGroupSum)
	}
}

type fakeCovidRepositoryFailure struct{}

func (f *fakeCovidRepositoryFailure) GetCovidData(req *entities.CovidDataReq) (*entities.CovidDataRes, error) {
	covidData := &entities.CovidDataRes{}
	return covidData, errors.New("Can not get data")
}

func TestSummaryCovidDataFailure(t *testing.T) {
	fakeCovidRepository := &fakeCovidRepositoryFailure{}
	covidUses := NewCovidUsecase(fakeCovidRepository)
	req := &entities.CovidDataReq{}
	_, err := covidUses.Summary(req)

	if err == nil {
		t.Errorf("expected failure but got success\n")
	}

	want := "Can not get data"
	if err := err.Error(); err != want {
		t.Errorf("%s is expected but got %s\n", want, err)
	}
}
