package api

import (
	"html"
	"net/http"
	"strings"

	validator "github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type Snipet struct {
	Id    uuid.UUID `json:"id"`
	UID   string    `json:"userId" valid:"required~userId->user id is needed"`
	Code  string    `json:"code"   valid:"required~code->it must be some code"`
	Title string    `json:"title"  valid:"required~title->title is needed,ascii~title->No special characters are allowed,stringlength(2|64)~title->Must be in between 2 to 64 characters"`
	Desc  string    `json:"desc"`
	Lang  string    `json:"lang"`
	Theme string    `json:"theme"`
}

type Response struct {
	Status int    `json:"status"`
	Error  string `json:"error,omitempty"`
	Data   Map    `json:"data"`
}

var (
	snipets      = make(map[uuid.UUID]Snipet)
	snipetsSlice = make([]Snipet, 0, 50)
)

func init() {
	app.GET("/snipets", func(ctx Context) {
		ctx.JSON(http.StatusOK, &Response{
			Status: http.StatusOK,
			Data: Map{
				"snipets":  snipetsSlice,
				"total":    len(snipetsSlice),
				"capacity": cap(snipetsSlice),
			},
		})
	})

	app.POST("/snipets", func(ctx Context) {

		snipet := Snipet{}
		snipet.Code = ctx.PostForm("code")
		snipet.Title = ctx.PostForm("title")
		snipet.Desc = ctx.PostForm("desc")
		snipet.Lang = ctx.PostForm("lang")
		snipet.Theme = ctx.PostForm("theme")
		snipet.UID = ctx.PostForm("userId")

		if snipet.UID == "" {
			ctx.JSON(400, &Response{
				400,
				"User id not provided",
				nil,
			})
			return
		}

		_, validUserId := users[snipet.UID]

		if !validUserId {
			ctx.JSON(401, &Response{
				401,
				"User does not exist",
				Map{
					"userId": snipet.UID,
				},
			})
			return
		}

		_, err := validator.ValidateStruct(snipet)

		errorFields := make([]Map, 0)

		if err != nil {
			errs := strings.Split(err.Error(), ";")

			for _, err := range errs {
				s := strings.Split(err, "->")
				errorFields = append(errorFields, Map{s[0]: s[1]})
			}

			ctx.JSON(http.StatusUnprocessableEntity, &Response{
				422,
				"Input validation error",
				Map{"fields": errorFields},
			})
			return
		}

		if ctx.PostForm("escaped") != "true" {
			snipet.Code = html.EscapeString(snipet.Code)
			snipet.Title = html.EscapeString(snipet.Title)
			if snipet.Desc != "" {
				snipet.Desc = html.EscapeString(snipet.Desc)
			}
		}

		if snipet.Lang == "" {
			snipet.Lang = "text"
		}

		if snipet.Theme == "" {
			snipet.Theme = "poimandres"
		}

		snipet.Id = uuid.New()
		snipets[snipet.Id] = snipet

		for _, value := range snipets {
			if value.Id == snipet.Id {
				snipetsSlice = append(snipetsSlice, value)
			}
		}

		ctx.JSON(200, &Response{
			Status: 200,
			Data: Map{
				"snipet": snipet,
			},
		})
	})
}

// export Handler type
func Snipets(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
