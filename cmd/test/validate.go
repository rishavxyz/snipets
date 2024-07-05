package test

import (
	"fmt"
	"strings"

	validator "github.com/asaskevich/govalidator"
)

type User struct {
	Id       string `json:"id" valid:"-"`
	Name     string `json:"name" valid:"required,ascii,stringlength(2|32)"`
	Username string `json:"username" valid:"required,alphanum,stringlength(4|64),matches(^[a-z][a-z0-9]+)~username->Only a-z and 0-9 allowed"`
	Password string `json:"password" valid:"required,stringlength(6|64)~password->From 6-64"`
}

func Validate() {
	user := User{
		Id:       "John Cena",
		Name:     "Rishav Mandal",
		Username: "1rishav",
		Password: "adj",
	}

	ok, err := validator.ValidateStruct(user)

	if err != nil {
		errs := strings.Split(err.Error(), ";")

		for _, err := range errs {
			s := strings.Split(err, "->")

			fmt.Printf("Field: %s\nError: %s\n", s[0], s[1])
		}

		return
	}

	fmt.Printf("User validation: %v\n", ok)
}
