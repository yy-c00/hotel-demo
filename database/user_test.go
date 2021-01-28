package database

import (
	"fmt"
	"testing"

	"github.com/yy-c00/hotel-demo/model"
)

var (
	manager = NewUserAccessManager()
	table []Test
)

type Test struct {
	model.User
	Err  error
	Expected bool
}

func init() {
	for i := 1; i < 5; i++ {
		data := fmt.Sprintf("Usuario%d", i)

		newUser := model.User{
			Name:     data,
			LastName: data,
			User:     data,
			Password: data,
			Root:     i%3 == 1,
		}

		newRow := Test{User: newUser, Expected: newUser.Root, Err: nil}

		table = append(table, newRow)
	}
}

func TestUserAccessManager_CreateNewUser(t *testing.T) {
	for _, row := range table {
		credential, err := manager.CreateNewUser(row.User)

		if err != row.Err || row.Expected != credential.Root {
			t.Errorf("Actual: (%v, %v), expected: (%v, %v)", err, credential.Root, row.Err, row.Expected)
		}

		t.Log(credential)
	}
}

func TestUserAccessManager_ValidateUser(t *testing.T) {
	for _, row := range table {
		user := row.User
		credential, err := manager.ValidateUser(model.Access{User: user.User, Password: user.Password})
		if err != row.Err || row.Expected != credential.Root {
			t.Logf("Data: %v", row.User)
			t.Errorf("Actual: (%v, %v), expected: (%v, %v)", err, credential.Root, row.Err, row.Expected)
		}
		t.Log(credential)
	}
}