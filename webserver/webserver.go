package webserver

import (
	"fmt"
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
			resp := ""
			for _, v := range timetable.Timetable[0].Events {
				resp += fmt.Sprintf("%s %s %s %s\n", v.Time, v.Name, v.Description, v.Place)
			}
			c.String(http.StatusOK, resp)
		}
	}
}

// CreateRoutes is creating group of routes
func CreateRoutes(ttWrapper *tt.Wrapper) func(g *gin.Engine) {
	return func(g *gin.Engine) {

		g.POST("/timetable", timetableHandler(ttWrapper))

		g.GET("/ping", ping)

	}
}
