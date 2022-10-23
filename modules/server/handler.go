package server

import (
	"github.com/gin-gonic/gin"
	_covidController "github.com/nuttapon-first/summarize-covid19-stats/modules/covid/controllers"
	_covidRepository "github.com/nuttapon-first/summarize-covid19-stats/modules/covid/repositories"
	_covidUseCase "github.com/nuttapon-first/summarize-covid19-stats/modules/covid/usecases"
)

func (s *Server) MapHandlers() error {

	covidPath := s.App.Group("/covid")
	covidRepository := _covidRepository.NewCovidRepository()
	covidUseCase := _covidUseCase.NewCovidUsecase(covidRepository)
	_covidController.NewCovidController(covidPath, covidUseCase)

	s.App.Use(func(c *gin.Context) {
		c.AbortWithStatusJSON(404, gin.H{
			"message": "end point not found",
		})
	})

	return nil
}
