package usecases

import (
	"sync"

	"github.com/nuttapon-first/summarize-covid19-stats/modules/entities"
)

type covidUses struct {
	CovidRepo entities.CovidRepository
}

func NewCovidUsecase(covidRepo entities.CovidRepository) entities.CovidUsecase {
	return &covidUses{
		CovidRepo: covidRepo,
	}
}

func (u *covidUses) Summary(req *entities.CovidDataReq) (*entities.CovidSummaryRes, error) {
	summarize := &entities.CovidSummaryRes{
		Province: make(map[string]int),
		AgeGroup: make(map[string]int),
	}

	data, err := u.CovidRepo.GetCovidData(req)
	if err != nil {
		return summarize, err
	}
	var wg sync.WaitGroup
	var ch = make(chan struct{}, 100) // limit go goroutine concurrent jobs
	for _, val := range data.Data {
		wg.Add(1)
		ch <- struct{}{}
		go func(data entities.CovidData) {
			defer wg.Done()
			summarize.CountData(data)
			<-ch
		}(val)
	}
	wg.Wait() // wait for all goroutine complete
	close(ch)

	return summarize, nil
}
