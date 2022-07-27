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

func timetableHandler(ttw *tt.Wrapper) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := c.BindJSON(ttw)
		if err == nil {
			c.Status(http.StatusNoContent)
		} else {
			println(err.Error())
		}
		ttw.FormatSelf()
	}
}

// CreateRoutes is creating group of routes
func CreateRoutes(ttWrapper *tt.Wrapper) func(g *gin.Engine) {
	return func(g *gin.Engine) {

		g.POST("/timetable", timetableHandler(ttWrapper))

		g.GET("/ping", ping)

	}
}
