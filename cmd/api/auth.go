package api

import (
	"fmt"
	"net/http"
)

func init() {
	App.GET("/auth", func(ctx CTX) {
		ctx.JSON(200, Map{"Helloo": "world"})
	})

	App.POST("/auth", func(ctx CTX) {
		var (
			name = ctx.PostForm("name")
		)

		fmt.Printf("name: %s\n", name)

		ctx.JSON(200, Map{"name": name})
	})
}

func Auth(w http.ResponseWriter, r *http.Request) {
	App.ServeHTTP(w, r)
}
