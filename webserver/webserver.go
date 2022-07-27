package webserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	tt "github.com/maxnrm/utro2022bot/timetable"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func timetableHandler(timetable *tt.Wrapper) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.BindJSON(&timetable) == nil {
			c.Status(http.StatusNoContent)
		}
		timetable.FormatSelf()
	}
}

// CreateRoutes is creating group of routes
func CreateRoutes(ttWrapper *tt.Wrapper) func(g *gin.Engine) {
	return func(g *gin.Engine) {

		g.POST("/timetable", timetableHandler(ttWrapper))

		g.GET("/ping", ping)

	}
}
