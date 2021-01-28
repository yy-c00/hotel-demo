package database

import (
	"fmt"

	"github.com/yy-c00/hotel-api/model"
	"golang.org/x/crypto/bcrypt"
)

type userAccessManager struct{}

//NewUserAccessManager returns an implementation of model.UserAccessManager
func NewUserAccessManager() model.UserAccessManager {
	return userAccessManager{}
}

func (u userAccessManager) UserIsLogged(credential model.Credential) error {
	if credential.Root {
		result := Connection().QueryRow("SELECT * FROM root WHERE idroot = ?", credential.ID)
		return result.Err()
	}

	result := Connection().QueryRow("SELECT * FROM employee WHERE idemployee = ?", credential.ID)
	return result.Err()
}

func (u userAccessManager) ValidateUser(access model.Access) (model.Credential, error) {
	var (
		id             uint
		credential     model.Credential
		hashedPassword string
	)

	row := Connection().QueryRow("SELECT iduser, name, lastname, password FROM user WHERE user = ?", access.User)

	err := row.Scan(&id, &credential.Name, &credential.LastName, &hashedPassword)
	if err != nil {
		return model.Credential{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(access.Password))
	if err != nil {
		return model.Credential{}, err
	}

	err = Connection().QueryRow("SELECT idroot FROM root WHERE iduser = ?", id).Scan(&credential.ID)
	if err == nil {
		credential.Root = true
		return credential, nil
	}

	err = Connection().QueryRow("SELECT idemployee FROM employee WHERE iduser = ?", id).Scan(&credential.ID)
	if err == nil {
		return credential, nil
	}

	return model.Credential{}, model.ErrAccessDenied
}

func (u userAccessManager) CreateNewUser(user model.User) (model.Credential, error) {
	tx, err := Connection().Begin()
	if err != nil {
		return model.Credential{}, err
	}

	insertIntoUser, err := tx.Prepare("INSERT INTO user(name, lastname, user, password) VALUES(?, ?, ?, ?)")
	if err != nil {
		return model.Credential{}, err
	}
	defer insertIntoUser.Close()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		tx.Rollback()
		return model.Credential{}, err
	}

	lastInsertUser, err := insertIntoUser.Exec(user.User, user.LastName, user.User, string(hashedPassword))
	if err != nil {
		return model.Credential{}, err
	}

	var table string
	if user.Root {
		table = "root"
	} else {
		table = "employee"
	}

	insertIntoTable, err := tx.Prepare(fmt.Sprintf("INSERT INTO %s(iduser) VALUES(?)", table))
	if err != nil {
		return model.Credential{}, err
	}
	defer insertIntoTable.Close()

	lastID, _ := lastInsertUser.LastInsertId()

	lastInsertTable, err := insertIntoTable.Exec(lastID)
	if err != nil {
		tx.Rollback()
		return model.Credential{}, err
	}

	lastID, _ = lastInsertTable.LastInsertId()

	credential := model.Credential{
		ID:       uint(lastID),
		Name:     user.Name,
		LastName: user.LastName,
		Root:     user.Root,
	}

	return credential, tx.Commit()
}

func (u userAccessManager) AvailableUsers() ([]model.Credential, error) {
	return nil, model.ErrFunctionIsNotDefined
}