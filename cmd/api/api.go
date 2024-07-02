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
	Snipets []Snipet = make([]Snipet, 0, 100)
	app     *gin.Engine
)

func init() {
	app = gin.New()

	app.GET("/snipets", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &Response{
			Status: http.StatusOK,
			Data: Data{
				unescape(&Snipets),
				len(Snipets),
				cap(Snipets),
			},
		})
	})

	app.POST("/snipets/new", func(ctx *gin.Context) {
		fields := make(map[string]string)

		fields["code"] = ctx.PostForm("code")
		fields["title"] = ctx.PostForm("title")
		fields["desc"] = ctx.PostForm("desc")
		fields["lang"] = ctx.PostForm("lang")
		fields["theme"] = ctx.PostForm("theme")

		emptyFields := make([]string, 0, len(fields))

		for key, value := range fields {
			switch key {
			case "lang":
				if value == "" {
					fields[key] = "text"
				}
			case "theme":
				if value == "" {
					fields[key] = "poimandres"
				}
			default:
				if key != "desc" && len(value) == 0 {
					emptyFields = append(emptyFields, key)
				} else {
					fields[key] = html.UnescapeString(value)
				}
			}
		}

		if len(emptyFields) > 0 {
			ctx.JSON(http.StatusBadRequest, &Response{
				Status: http.StatusBadRequest,
				Error:  "Fields cannot be empty",
				Data: gin.H{
					"emptyFields": emptyFields,
				},
			})
			return
		}

		snipet := Snipet{
			Id:    uuid.New(),
			Code:  fields["code"],
			Title: fields["title"],
			Desc:  fields["desc"],
			Lang:  fields["lang"],
			Theme: fields["theme"],
		}

		Snipets = append(Snipets, snipet)

		ctx.JSON(http.StatusOK, &Response{
			Status: http.StatusOK,
			Data:   gin.H{"snipet": &snipet},
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
