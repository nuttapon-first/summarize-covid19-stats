package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nuttapon-first/summarize-covid19-stats/modules/entities"
)

type response struct {
	ErrMessage string `json:"message"`
}

type fakeCovidUse struct{}

func (f *fakeCovidUse) Summary(req *entities.CovidDataReq) (*entities.CovidSummaryRes, error) {
	covidSummaryData := &entities.CovidSummaryRes{}
	return covidSummaryData, errors.New("something went wrong")
}

func TestSummaryCovidFailure(t *testing.T) {
	fakeCovidUseCase := &fakeCovidUse{}
	fakeController := &covidController{CovidUse: fakeCovidUseCase}

	res := httptest.NewRecorder()
	c, r := gin.CreateTestContext(res)
	c.Request = httptest.NewRequest(http.MethodGet, "/covid/summary", nil)
	c.Request.Header.Set("Content-Type", "application/json; charset=utf-8")

	r.GET("/covid/summary", fakeController.Summary) // Call to a handler method
	r.ServeHTTP(res, c.Request)

	if status := res.Code; status != 500 {
		t.Errorf("handler returned wrong status code: got %v want %v", status, 500)
	}

	response := &response{}
	json.Unmarshal(res.Body.Bytes(), response)

	want := "something went wrong"

	if got := response.ErrMessage; got != want {
		t.Errorf("Message: got %v want %v", got, want)
	}
}

type fakeCovidUseSuccess struct{}

func (f *fakeCovidUseSuccess) Summary(req *entities.CovidDataReq) (*entities.CovidSummaryRes, error) {
	covidSummaryData := &entities.CovidSummaryRes{
		Province: map[string]int{
			"Bangkok": 1,
		},
		AgeGroup: map[string]int{
			"0-30":  0,
			"31-60": 0,
			"61+":   0,
			"N/A":   1,
		},
	}
	return covidSummaryData, nil
}

func TestSummaryCovidSuccess(t *testing.T) {
	fakeCovidUseCase := &fakeCovidUseSuccess{}
	fakeController := &covidController{CovidUse: fakeCovidUseCase}

	res := httptest.NewRecorder()
	c, r := gin.CreateTestContext(res)
	c.Request = httptest.NewRequest(http.MethodGet, "/covid/summary", nil)
	c.Request.Header.Set("Content-Type", "application/json; charset=utf-8")

	r.GET("/covid/summary", fakeController.Summary) // Call to a handler method
	r.ServeHTTP(res, c.Request)

	if status := res.Code; status != 200 {
		t.Errorf("handler returned wrong status code: got %v want %v", status, 200)
	}

	response := &entities.CovidSummaryRes{}
	json.Unmarshal(res.Body.Bytes(), response)

	wantBangkok := 1
	if got := response.Province["Bangkok"]; got != wantBangkok {
		t.Errorf("Message: got %v want %v", got, wantBangkok)
	}

	wantNA := 1
	if got := response.AgeGroup["N/A"]; got != wantNA {
		t.Errorf("Message: got %v want %v", got, wantNA)
	}
}
