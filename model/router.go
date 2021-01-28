package model

import "github.com/labstack/echo/v4"

const (
	//NewSale is used to report that create a new sale
	NewSale uint = iota + 1
	//ConfirmedSale is used to report that a sale was confirmed
	ConfirmedSale
	//ConfirmedSale is used to report that a sale was canceled
	CanceledSale
)

//Message used to send as websocket message
type Message struct {
	//Type indicate type of message
	Type uint        `json:"type"`
	//Data data of message
	Data interface{} `json:"data"`
}

//Router contains a group routers
//(model.SalesRouter, model.StorageRouter, model.HistoryRouter, model.UserAccessRouter, model.RoomAccessRouter)
type Router struct {
	SalesRouter
	StorageRouter
	HistoryRouter
	UserAccessRouter
	RoomAccessRouter
}

//UserAccessRouter interface implemented by all structures to manage user access routes
type UserAccessRouter interface {
	//UserIsLogged echo.HandleFunc to validate if user is logged
	UserIsLogged(echo.Context) error
	//ValidateUser echo.HandleFunc to login (validate if user is valid)
	ValidateUser(echo.Context) error
	//CreateNewUser echo.HandleFunc to create a new user
	CreateNewUser(echo.Context) error
	//AvailableUsers echo.HandleFunc to get all existing users
	AvailableUsers(echo.Context) error
}

//RoomAccessRouter interface implemented by all structures to manage room access routes
type RoomAccessRouter interface {
	//RoomIsLogged echo.HandleFunc to validate if room is logged
	RoomIsLogged(echo.Context) error
	//ValidateRoom echo.HandleFunc to validate if room exists
	ValidateRoom(echo.Context) error
	//CreateNewRoom echo.HandleFunc to create a new room
	CreateNewRoom(echo.Context) error
	//AvailableRooms echo.HandleFunc to get all existing rooms
	AvailableRooms(echo.Context) error
}

//SalesRouter interface implemented by all structures to manage sales routes
type SalesRouter interface {
	//NewSale echo.HandleFunc to make a new sale
	NewSale(echo.Context) error
	//ConfirmSale echo.HandleFunc to confirm an existing sale
	ConfirmSale(echo.Context) error
	//CancelSale echo.HandleFunc to cancel/delete an existing sale
	CancelSale(echo.Context) error
	//GetAllPendingSales echo.HandleFunc to get all pending sales
	GetAllPendingSales(echo.Context) error
	//GetSaleById echo.HandleFunc to get a sale by id
	GetSaleById(echo.Context) error
	//Channel echo.HandleFunc to make a handshake with the sales channel (ws)
	Channel(echo.Context) error
}

//StorageRouter interface implemented by all structures to manage storage routes
type StorageRouter interface {
	//AddProduct echo.HandleFunc to add product
	AddProduct(echo.Context) error
	//SetProduct echo.HandleFunc to update a existing product
	SetProduct(echo.Context) error
	//DeleteProduct echo.HandleFunc to delete a existing product
	DeleteProduct(echo.Context) error

	//AddCategory echo.HandleFunc to add category to storage
	AddCategory(echo.Context) error
	//AddCategories echo.HandleFunc to add various categories to storage
	AddCategories(echo.Context) error
	//DeleteCategory echo.HandleFunc to delete a some category by id
	DeleteCategory(echo.Context) error

	//GetAll echo.HandleFunc to get all categories from storage
	GetAll(echo.Context) error
	//GetProductByID echo.HandelFunc to get some product by id
	GetProductByID(echo.Context) error
	//GetCategoryByID echo.HandleFunc to get some category by id
	GetCategoryByID(echo.Context) error

	//Log echo.HandleFunc to add and log a new quantity to be added
	Log(echo.Context) error
	//Search echo.HandleFunc to add and log a new quantity of some product
	Search(echo.Context) error
	//Channel echo.HandleFunc to make a handshake with the storage channel (ws)
	Channel(echo.Context) error
}

//HistoryRouter interface implemented by all structures to manage history routes
type HistoryRouter interface {
	//GetLogsByRange echo.HandleFunc to get all confirmed sales by range
	GetLogsByRange(echo.Context) error
	//GetSaleLogById echo.HandleFunc to get a log by id
	GetSaleLogById(echo.Context) error
	//SearchByRange echo.HandleFunc to search history by range
	SearchByRange(echo.Context) error
}
