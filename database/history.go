package database

import "github.com/yy-c00/hotel-api/model"

type history struct {}

func NewHistory() model.History {
	return history{}
}

func (h history)  GetLogsByRange(left, right uint) ([]model.Sale, error) {
	sales := []model.Sale{}
	results, err := Connection().Query("SELECT idsales FROM sales WHERE status = TRUE ORDER BY time_stamp DESC LIMIT ?, ?", left, right)
	if err != nil {
		return nil, err
	}

	for results.Next() {
		var idSale uint

		err = results.Scan(&idSale)
		if err != nil {
			return nil, err
		}

		sale, err := h.GetSaleLogById(idSale)
		if err != nil {
			return nil, err
		}

		sales = append(sales, sale)
	}

	return sales, nil
}

func (history) GetSaleLogById(id uint) (model.Sale, error){
	sale := model.Sale{ID: id, Room: model.Room{}}

	err := Connection().QueryRow("SELECT status, time_stamp FROM sales WHERE idsales = ?", sale.ID).Scan(&sale.Status, &sale.Date)
	if err != nil {
		return model.Sale{}, err
	}

	Connection().QueryRow("CALL SEARCH_ROOM_SALE(?)", sale.ID).Scan(&sale.Room.ID)

	user := model.Credential{}

	err = Connection().QueryRow("CALL SEARCH_EMPLOYEE_SALE(?)", sale.ID).Scan(&user.ID, &user.Name, &user.LastName)
	if err != nil {

		err := Connection().QueryRow("CALL SEARCH_ROOT_SALE(?)", sale.ID).Scan(&user.ID, &user.Name, &user.LastName)

		if err != nil {
			return model.Sale{}, err
		}

		user.Root = true
	}

	sale.Credential = user

	return sale, nil
}


func (h history) SearchByRange(str string, left, right uint) ([]model.Sale, error) {
	sales := []model.Sale{}

	results, err := Connection().Query("SELECT idsales FROM sales WHERE status = TRUE AND time_stamp LIKE ? LIMIT ?, ?", "%" + str + "%", left, right)
	if err != nil {
		return nil, err
	}

	for results.Next() {
		sale := model.Sale{}

		err = results.Scan(&sale.ID)
		if err != nil {
			return nil, err
		}

		sale, err = h.GetSaleLogById(sale.ID)
		if err != nil {
			return nil, err
		}

		sales = append(sales, sale)
	}

	return sales, nil
}