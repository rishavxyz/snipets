package api

import (
	"html"
	"net/http"

	gin "snipets/app"

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

type Response struct {
	Status int     `json:"status"`
	Error  string  `json:"error,omitempty"`
	Data   gin.Map `json:"data"`
}

var (
	Snipets []Snipet = make([]Snipet, 0, 100)
)

func init() {
	gin.App.GET("/snipets", func(ctx gin.CTX) {
		ctx.JSON(http.StatusOK, &Response{
			Status: http.StatusOK,
			Data: gin.Map{
				"snipets":  unescape(&Snipets),
				"total":    len(Snipets),
				"capacity": cap(Snipets),
			},
		})
	})

	gin.App.POST("/snipets/new", func(ctx gin.CTX) {
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
				Data: gin.Map{
					"emptyFields": emptyFields,
				},
			})
			return
		}

		snipet := &Snipet{
			Id:    uuid.New(),
			Code:  fields["code"],
			Title: fields["title"],
			Desc:  fields["desc"],
			Lang:  fields["lang"],
			Theme: fields["theme"],
		}

		Snipets = append(Snipets, *snipet)

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
func Main(w http.ResponseWriter, r *http.Request) {
	gin.App.ServeHTTP(w, r)
}
