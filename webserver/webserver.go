package webserver

import (
	"github.com/gin-gonic/gin"
	"github.com/maxnrm/utro2022bot/db"
	tt "github.com/maxnrm/utro2022bot/timetable"
)

var dbHandler db.Handler = db.DBHandler

// WebServer exported global webserver (gin) instance
var WebServer *gin.Engine = New()

// New create new webserver instance with added route handlers
func New() *gin.Engine {
	r := gin.Default()

	r.SetTrustedProxies(nil)

	createRoutes(&ttWrapper)(r)

	return r
}

func createRoutes(ttWrapper *tt.Wrapper) func(g *gin.Engine) {
	return func(g *gin.Engine) {
		g.POST("/timetable", timetableHandler)
		g.GET("/ping", ping)
	}
}
