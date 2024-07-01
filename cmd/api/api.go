package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Snipet struct {
	Id    uuid.UUID `json:"id"`
	Code  string    `json:"code"`
	Title string    `json:"title"`
	Desc  string    `json:"desc"`
	Lang  string    `json:"lang"`
	Theme string    `json:"theme"`
}

type Data struct {
	Snipets  []*Snipet `json:"snipets"`
	Length   int       `json:"length"`
	Capacity int       `json:"capacity"`
}

type Response struct {
	Status  int   `json:"status"`
	Error   error `json:"error"`
	Success bool  `json:"success"`
	Data    Data  `json:"data"`
}

var (
	snipets []*Snipet = make([]*Snipet, 0, 100)
	app     *gin.Engine
)

func init() {
	snipets = append(snipets,
		&Snipet{
			uuid.New(),
			"const name = \"Rishav\"",
			"This is JavaScript",
			"",
			"javascript",
			"nord",
		},
		&Snipet{
			uuid.New(),
			":(){:|:&};:",
			"Fork bomb (Dos attack)",
			"Do NOT run this code, it will crash your system",
			"bash",
			"dracula",
		},
		&Snipet{
			uuid.New(),
			"void main() {\n  println(\"Welcome to Kt\")\n}",
			"Some kotlin",
			"Kotlin is fun",
			"kotlin",
			"catppuccin-mocha",
		},
	)

	app = gin.New()
	r := app.Group("/api")

	postNewSnipet(r)
	getSnipets(r)
}

// export Handler type
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}

// endpoints
func postNewSnipet(r *gin.RouterGroup) {
	r.POST("/new", func(ctx *gin.Context) {
		var (
			code  = ctx.PostForm("code")
			title = ctx.PostForm("title")
			desc  = ctx.PostForm("desc")
			lang  = ctx.PostForm("lang")
			theme = ctx.PostForm("theme")
		)

		snipets = append(snipets, &Snipet{
			uuid.New(), code, title, desc, lang, theme,
		})

		ctx.JSON(http.StatusOK, Response{
			http.StatusOK, nil, true, Data{
				snipets,
				len(snipets),
				cap(snipets),
			},
		})
	})
}

func getSnipets(r *gin.RouterGroup) {
	r.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, Response{
			http.StatusOK, nil, true, Data{
				snipets,
				len(snipets),
				cap(snipets),
			},
		})
	})
}
