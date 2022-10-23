package repositories

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/nuttapon-first/summarize-covid19-stats/modules/entities"
)

type covidRepo struct{}

func NewCovidRepository() entities.CovidRepository {
	return &covidRepo{}
}

func (r *covidRepo) GetCovidData(req *entities.CovidDataReq) (*entities.CovidDataRes, error) {
	covidData := &entities.CovidDataRes{}
	data, err := GetDataFromStaticUrl()

	if err != nil {
		log.Println(err)
		return covidData, err
	}

	json.Unmarshal([]byte(data), covidData)
	return covidData, nil
}

func GetDataFromStaticUrl() (string, error) {
	url := os.Getenv("COVID_CASE_URL")
	conn, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return "", err
	}

	defer conn.Body.Close()
	body, err := io.ReadAll(conn.Body)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	covidData := string(body)
	return covidData, nil
}
