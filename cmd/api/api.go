package api

import (
	"html"
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
	Snipets  *[]Snipet `json:"snipets"`
	Length   int       `json:"length"`
	Capacity int       `json:"capacity"`
}

type Response struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
	Data   any    `json:"data"`
}

var (
	Snipets = make([]Snipet, 0, 100)
	app     *gin.Engine
)

func init() {
	app = gin.New()
	r := app.Group("/api")

	r.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &Response{
			Status: http.StatusOK,
			Data: Data{
				unescape(&Snipets),
				len(Snipets),
				cap(Snipets),
			},
		})
	})

	r.POST("/new", func(ctx *gin.Context) {
		var (
			code  = ctx.PostForm("code")
			title = ctx.PostForm("title")
			desc  = ctx.PostForm("desc")
			lang  = ctx.PostForm("lang")
			theme = ctx.PostForm("theme")
		)

		if code == "" || title == "" {
			ctx.JSON(http.StatusBadRequest, &Response{
				Status: http.StatusBadRequest,
				Error:  "Code and title cannot be empty",
				Data: gin.H{
					"code":  code,
					"title": title,
				},
			})
			return
		}

		snipet := Snipet{
			uuid.New(), code, title, desc, lang, theme,
		}

		Snipets = append(Snipets, snipet)

		ctx.JSON(http.StatusOK, &Response{
			Status: http.StatusOK,
			Data:   gin.H{"snipet": snipet},
		})
	})
}

func unescape(snipets *[]Snipet) *[]Snipet {
	for _, v := range *snipets {
		v.Code = html.UnescapeString(v.Code)
		v.Title = html.UnescapeString(v.Title)
		v.Desc = html.UnescapeString(v.Desc)
	}
	return snipets
}

// export Handler type
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
