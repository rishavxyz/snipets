package api

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"html"
	"net/http"
	"strings"

	validator "github.com/asaskevich/govalidator"
)

type UserID = string

type User struct {
	Id       UserID `json:"id"                 valid:"-"`
	Name     string `json:"name"               valid:"required~name->name is needed,ascii~name->No special characters allowrd,stringlength(2|32)~name->Must be in between 2 to 32 characters"`
	Username string `json:"username"           valid:"required~username->username is needed,alphanum~username->Only A-Z and 0-9 allowed,stringlength(4|64)~username->Must be in between 4 to 64 characters,matches(^[a-z][a-z0-9]+)~username->Cannot start with a number"`
	Password string `json:"password,omitempty" valid:"required~password->password is needed,stringlength(6|64)~password->Must be in between 6 to 64 characters"`
}

var (
	users      = make(map[UserID]User)
	usersSlice = make([]User, 0, 50)
)

func init() {
	app.GET("/users", func(ctx Context) {
		ctx.JSON(http.StatusOK, &Response{
			Status: 200,
			Data: Map{
				"users":    usersSlice,
				"total":    len(usersSlice),
				"capacity": cap(usersSlice),
			},
		})
	})

	app.POST("/users", func(ctx Context) {
		user := User{}
		user.Name = ctx.PostForm("name")
		user.Username = ctx.PostForm("uname")
		user.Password = ctx.PostForm("passwd")

		_, err := validator.ValidateStruct(user)

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

		user.Id = CreateUserId(&user.Username)

		if _, found := users[user.Id]; found {
			ctx.JSON(http.StatusConflict, &Response{
				409,
				"Username already taken",
				Map{"username": user.Username},
			})
			return
		}

		users[user.Id] = user

		if ctx.PostForm("hashed") != "true" {
			user.Password = hash(&user.Password)
		}
		if ctx.PostForm("escaped") != "true" {
			user.Name = html.EscapeString(user.Name)
			user.Username = html.EscapeString(user.Username)
		}

		for _, value := range users {
			if value.Id == user.Id {
				usersSlice = append(usersSlice, value)
			}
		}

		ctx.JSON(200, &Response{
			Status: 200,
			Data: Map{
				"user": User{
					Id:       user.Id,
					Name:     user.Name,
					Username: user.Username,
				},
			},
		})
	})
}

func Users(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}

func hash(password *string) string {
	hash := sha512.New()
	hash.Write([]byte(*password))
	return hex.EncodeToString(hash.Sum(nil))
}

func CreateUserId(username *string) string {
	hash := md5.New()
	hash.Write([]byte(*username))
	return hex.EncodeToString(hash.Sum(nil))
}
