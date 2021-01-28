package router

import (
	"github.com/yy-c00/hotel-demo/model"
	"net/http"
	"testing"
)

var (
	userRouter = NewUserAccessRouter()
	userTable = []RequestConfig{
		{
			"CreateNewUser",
			"",
			&model.User{Name: "y", LastName: "y", User: "another user", Password: "strong"},
			http.MethodPost,
			userRouter.CreateNewUser,
			http.StatusCreated,
		},
		{
			"ValidateUser",
			"",
			&model.Access{User: "another user", Password: "strong"},
			http.MethodPost,
			userRouter.ValidateUser,
			http.StatusOK,
		},
		{
			"AvailableUsers",
			"",
			nil,
			http.MethodGet,
			userRouter.AvailableUsers,
			http.StatusOK,
		},
	}
)

func TestAccessRouter(t *testing.T) {
	ConfigTest(userTable, t)
}
