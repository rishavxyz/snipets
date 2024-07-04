package api

import (
	"crypto/sha256"
	"html"
	"net/http"

	gin "snipets/app"

	"github.com/google/uuid"
)

type User struct {
	Id       string
	Name     string
	Username string
	Password string
}

func init() {
	gin.App.GET("/users") // TODO

	gin.App.POST("/users/new", func(ctx gin.CTX) {
		user := make(map[string]string)
		user["name"] = ctx.PostForm("name")
		user["username"] = ctx.PostForm("uname")
		user["password"] = ctx.PostForm("passwd")

		emptyFields := make([]string, 0, len(user))

		for key, value := range user {
			switch key {
			case "password":
				if value == "" {
					emptyFields = append(emptyFields, key)
				} else if len(value) < 8 || len(value) > 28 {
					ctx.JSON(406, &Response{
						Status: 406,
						Error:  "Password must be between 8 to 28 letters",
					})
					return
				} else {
					if ctx.PostForm("hashed") != "true" {
						h := sha256.New()
						h.Write([]byte(value))
						user["password"] = string(h.Sum(nil))
					}
				}
			default:
				if value == "" {
					emptyFields = append(emptyFields, key)
				} else {
					user[key] = html.EscapeString(value)
				}
			}
		}

		if len(emptyFields) > 0 {
			ctx.JSON(400, &Response{
				Status: 400,
				Error:  "Username and password are must",
				Data: gin.Map{
					"emptyFields": emptyFields,
				},
			})
			return
		}
		user["id"] = uuid.New().String()

		users := make([]User, 0, 50)
		newUser := &User{
			Id:       user["id"],
			Name:     user["name"],
			Username: user["username"],
			Password: user["password"],
		}
		users = append(users, *newUser)

		ctx.JSON(200, &Response{
			Status: 200,
			Data: gin.Map{
				"userCreated": true,
				"newUser":     *newUser,
				"users":       users,
			},
		})
	})
}

func Users(w http.ResponseWriter, r *http.Request) {
	gin.App.ServeHTTP(w, r)
}
