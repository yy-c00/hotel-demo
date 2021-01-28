package database

import "github.com/yy-c00/hotel-api/model"

type roomAccessManager struct{}

//NewRoomAccessManager returns an implementation of model.RoomAccessManager
func NewRoomAccessManager() model.RoomAccessManager {
	return roomAccessManager{}
}

func (r roomAccessManager) RoomIsLogged(room model.Room) error {
	result := Connection().QueryRow("SELECT * FROM room WHERE idroom = ?", room.ID)
	return result.Err()
}

func (r roomAccessManager) ValidateRoom(room model.Room) error {
	err := Connection().QueryRow("SELECT * FROM room WHERE idroom = ?", room.ID).Scan(&room.ID)
	return err
}

func (r roomAccessManager) CreateNewRoom(room model.Room) error {
	_, err := Connection().Exec("INSERT INTO room values(?)", room.ID)
	return err
}

func (r roomAccessManager) AvailableRooms() ([]model.Room, error) {
	var rooms []model.Room

	results, err := Connection().Query("SELECT * FROM room")
	if err != nil {
		return nil, err
	}

	for results.Next() {
		var newRoom model.Room

		err = results.Scan(&newRoom.ID)
		if err != nil {
			return nil, err
		}

		rooms = append(rooms, newRoom)
	}

	return rooms, nil
}