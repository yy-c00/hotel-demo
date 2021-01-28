package database

import (
	"github.com/yy-c00/hotel-api/model"
	"testing"
)

var (
	roomManager  = NewRoomAccessManager()
	rooms = *new([]model.Room)
)

func init() {
	for i := uint(1); i <= 10; i++ {
		rooms = append(rooms, model.Room{ID: i})
	}
}

func TestRoomAccessManager_CreateNewRoom(t *testing.T) {
	for i := 0; i < len(rooms); i++ {
		err := roomManager.CreateNewRoom(rooms[i])
		if err != nil {
			t.Errorf("Occurs an unexpected error %s", err.Error())
		}
	}
}

func TestRoomAccessManager_AvailableRooms(t *testing.T) {
	arr, err := roomManager.AvailableRooms()
	if err != nil {
		t.Errorf("Occurs an unexpected error: %v", err.Error())
	}

	for i := 0; i < len(rooms); i++ {
		if arr[i] != rooms[i] {
			t.Errorf("Expeted: %v Got: %v", rooms[i], arr[i])
		}
	}
}

func TestRoomAccessManager_ValidateRoom(t *testing.T) {
	for i := 0; i < len(rooms); i++ {
		err := roomManager.ValidateRoom(rooms[i])
		if err != nil {
			t.Errorf("Room in not valid (%v)", err.Error())
		}
	}
}