package webserver

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	tt "github.com/maxnrm/utro2022bot/timetable"
)

var ttWrapper tt.Wrapper = tt.TimetableWrapper

// there route handlers goes
var (
	timetableHandler gin.HandlerFunc = timetableHandlerFactory(&ttWrapper)
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func timetableHandlerFactory(ttw *tt.Wrapper) gin.HandlerFunc {
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
