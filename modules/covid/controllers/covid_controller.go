package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nuttapon-first/summarize-covid19-stats/modules/entities"
)

type covidController struct {
	CovidUse entities.CovidUsecase
}

func NewCovidController(r *gin.RouterGroup, covidUse entities.CovidUsecase) {
	controllers := &covidController{
		CovidUse: covidUse,
	}
	r.GET("/summary", controllers.Summary)
}

func (h *covidController) Summary(c *gin.Context) {
	req := new(entities.CovidDataReq)

	res, err := h.CovidUse.Summary(req)

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"Province": res.Province,
		"AgeGroup": res.AgeGroup,
	})
}
