package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine = gin.New()

type Context = *gin.Context
type Map = map[string]interface{}

func init() {
	app.GET("/auth", func(ctx Context) {
		ctx.JSON(200, Map{"Helloo": "world"})
	})

	app.POST("/auth", func(ctx Context) {
		var (
			name = ctx.PostForm("name")
		)

		fmt.Printf("name: %s\n", name)

		ctx.JSON(200, Map{"name": name})
	})
}

func Auth(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
