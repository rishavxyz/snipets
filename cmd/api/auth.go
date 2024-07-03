package api

import (
	"fmt"
	"net/http"
	gin "snipets/app"
)

func init() {
	gin.App.GET("/auth", func(ctx gin.CTX) {
		ctx.JSON(200, gin.Map{"Helloo": "world"})
	})

	gin.App.POST("/auth", func(ctx gin.CTX) {
		var (
			name = ctx.PostForm("name")
		)

		fmt.Printf("name: %s\n", name)

		ctx.JSON(200, gin.Map{"name": name})
	})
}

func Auth(w http.ResponseWriter, r *http.Request) {
	gin.App.ServeHTTP(w, r)
}
