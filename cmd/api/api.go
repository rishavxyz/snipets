package api

import (
	"errors"
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
	Snipets  []*Snipet `json:"snipets"`
	Length   int       `json:"length"`
	Capacity int       `json:"capacity"`
}

type Response struct {
	Status int   `json:"status"`
	Error  error `json:"error"`
	Data   any   `json:"data"`
}

var (
	Snipets []*Snipet = make([]*Snipet, 0, 100)
	app     *gin.Engine
)

func init() {
	app = gin.New()

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, response(http.StatusOK, nil))
	})

	app.POST("/new", func(ctx *gin.Context) {
		var (
			code  = ctx.PostForm("code")
			title = ctx.PostForm("title")
			desc  = ctx.PostForm("desc")
			lang  = ctx.PostForm("lang")
			theme = ctx.PostForm("theme")
		)

		if code == "" || title == "" {
			ctx.JSON(http.StatusBadRequest, &Response{
				http.StatusBadRequest,
				errors.New("code and title cannot be empty"),
				gin.H{
					"code":  code,
					"title": title,
				},
			})
		}

		snipet := &Snipet{
			uuid.New(), code, title, desc, lang, theme,
		}

		Snipets = append(Snipets, snipet)

		ctx.JSON(http.StatusOK, &Response{
			http.StatusOK,
			nil,
			gin.H{"snipet": snipet},
		})
	})
}

func response(status int, err error, raw ...bool) *Response {
	if !(len(raw) == 1 && raw[0]) {
		for _, v := range Snipets {
			v.Title = html.UnescapeString(v.Title)
			v.Desc = html.UnescapeString(v.Desc)
			v.Code = html.UnescapeString(v.Code)
		}
	}
	return &Response{
		status,
		err,
		Data{
			Snipets,
			len(Snipets),
			cap(Snipets),
		},
	}
}

// export Handler type
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
