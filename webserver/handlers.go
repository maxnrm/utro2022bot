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
		body, err := ioutil.ReadAll(c.Request.Body)
		json.Compact(compactBody, body)

		dbHandler.AddTimetable(string(compactBody.Bytes()))

		if err == nil {
			c.Status(http.StatusNoContent)
			ttwt := tt.New()
			for i := range ttw.Timetables {
				ttw.Timetables[i] = tt.Timetable{}
			}
			for i, v := range ttwt.Timetables {
				ttw.Timetables[i] = v
			}
		} else {
			println(err.Error())
		}
	}
}
