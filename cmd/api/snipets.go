package api

import (
	"html"
	"net/http"

	"github.com/google/uuid"
)

type Snipet struct {
	Id    string `json:"id"`
	Code  string `json:"code"  valid:"required"`
	Title string `json:"title" valid:"required,stringlength(2|64)"`
	Desc  string `json:"desc"`
	Lang  string `json:"lang"  valid:"required"`
	Theme string `json:"theme"`
}

type Response struct {
	Status int    `json:"status"`
	Error  string `json:"error,omitempty"`
	Data   Map    `json:"data"`
}

var (
	snipets []Snipet = make([]Snipet, 0, 50)
)

func init() {
	app.GET("/snipets", func(ctx Context) {
		ctx.JSON(http.StatusOK, &Response{
			Status: http.StatusOK,
			Data: Map{
				"snipets":  unescape(&snipets),
				"total":    len(snipets),
				"capacity": cap(snipets),
			},
		})
	})

	app.POST("/snipets", func(ctx Context) {
		isEscaped := ctx.PostForm("escaped")

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
					if isEscaped == "" {
						fields[key] = html.EscapeString(value)
					} else {
						fields[key] = value
					}
				}
			}
		}

		if len(emptyFields) > 0 {
			ctx.JSON(http.StatusBadRequest, &Response{
				Status: http.StatusBadRequest,
				Error:  "Fields cannot be empty",
				Data: Map{
					"emptyFields": emptyFields,
				},
			})
			return
		}

		snipet := &Snipet{
			Id:    uuid.NewString(),
			Code:  fields["code"],
			Title: fields["title"],
			Desc:  fields["desc"],
			Lang:  fields["lang"],
			Theme: fields["theme"],
		}

		snipets = append(snipets, *snipet)

		ctx.JSON(http.StatusOK, &Response{
			Status: http.StatusOK,
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
func Snipets(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
