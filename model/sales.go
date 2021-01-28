package model

import "errors"

//ErrFailedSale returns when trying to make a sale and a unexpected error occurs
var ErrFailedSale = errors.New("failed sale")

//Sold sold
type Sold struct {
	Product  uint `json:"product"`
	Quantity uint `json:"quantity"`
	Sale     uint `json:"sale"`
}

//Sale sale
type Sale struct {
	ID         uint       `json:"id"`
	Date       string     `json:"date,omitempty"`
	Room       Room       `json:"room"`
	Status     bool       `json:"status"`
	Credential Credential `json:"credential"`
	Products   []Product  `json:"products,omitempty"`
}

//SalesManager interface that implemented all structures to manage sales
type SalesManager interface {
	//NewSale create a new sale
	NewSale(*Sale) error
	//ConfirmSale used to confirm a previous created sale
	ConfirmSale(Sale) error
	//CancelSale used to cancel/delete a existing sale
	CancelSale(uint) error
	//GetAllPendingSales used to get all pending sales
	GetAllPendingSales() ([]Sale, error)
	//GetSaleById used to get a sale by id
	GetSaleById(uint) (Sale, error)
}
