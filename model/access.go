package model

import (
	"errors"
)

var (
	//ErrAccessDenied returns when the user try to use a
	ErrAccessDenied error = errors.New("you do not have access")
	//ErrFunctionIsNotDefined returns or throws when a some function is undefined
	ErrFunctionIsNotDefined error = errors.New("this function is not defined")
)

//Room room
type Room struct {
	//ID room id
	ID uint `json:"id"`
}

//User user
type User struct {
	//Name user name
	Name     string `json:"name"`
	//LastName user lastname
	LastName string `json:"lastname"`
	//User user
	User     string `json:"user"`
	//Password password
	Password string `json:"password"`
	//Root specify if user is root
	Root     bool   `json:"root"`
}

//Access struct used to validate an user's access
type Access struct {
	//User user
	User     string `json:"user"`
	//Password password
	Password string `json:"password"`
}

//Credential struct used to identify a user
type Credential struct {
	//ID user id
	ID       uint   `json:"id"`
	//Name user name
	Name     string `json:"name"`
	//LastName user lastname
	LastName string `json:"lastname"`
	//Root specify if user is root
	Root     bool   `json:"root"`
}

//UserAccessManager interface implemented by all structs to manage user access
type UserAccessManager interface {
	//UserIsLogged validate if an user is logged
	UserIsLogged(Credential) error
	//ValidateUser verify that a user has access
	ValidateUser(Access) (Credential, error)
	//CreateNewUser create a new user
	CreateNewUser(User) (Credential, error)
	//AvailableUsers returns all existing users
	AvailableUsers() ([]Credential, error)
}

//UserAccessManager interface implemented by all structs to manage room access
type RoomAccessManager interface {
	//RoomIsLogged validate that a room is logged
	RoomIsLogged(Room) error
	//ValidateRoom verify that a room has access
	ValidateRoom(Room) error
	//CreateNewRoom create a new room in the database
	CreateNewRoom(Room) error
	//AvailableRooms returns all existing rooms
	AvailableRooms() ([]Room, error)
}
