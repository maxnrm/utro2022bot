package webserver

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxnrm/utro2022bot/db"
	tt "github.com/maxnrm/utro2022bot/timetable"
)

var dbHandler db.Handler = db.DBHandler

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func timetableHandler(ttw *tt.Wrapper) gin.HandlerFunc {
	return func(c *gin.Context) {
		var compactBody *bytes.Buffer = new(bytes.Buffer)
		body, _ := ioutil.ReadAll(c.Request.Body)
		json.Compact(compactBody, body)

		err := json.Unmarshal(compactBody.Bytes(), &ttw)
		if err == nil {
			c.Status(http.StatusNoContent)
		} else {
			println(err.Error())
		}
		ttw.FormatSelf()

		dbHandler.AddTimetable(string(compactBody.Bytes()))
	}
}

// CreateRoutes is creating group of routes
func CreateRoutes(ttWrapper *tt.Wrapper) func(g *gin.Engine) {
	return func(g *gin.Engine) {

		g.POST("/timetable", timetableHandler(ttWrapper))

		g.GET("/ping", ping)

	}
}
