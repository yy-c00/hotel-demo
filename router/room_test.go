package router

import (
	"github.com/yy-c00/hotel-demo/model"
	"net/http"
	"testing"
)

var (
	roomRouter = NewRoomAccessRouter()
	roomTable = []RequestConfig {
		{
			"CreateNewRoom",
			"",
			model.Room{ID: 1},
			http.MethodPost,
			roomRouter.CreateNewRoom,
			http.StatusCreated,
		},
		{
			"ValidateRoom",
			"",
			model.Room{ID: 1},
			http.MethodPost,
			roomRouter.ValidateRoom,
			http.StatusOK,
		},
		{
			"AvailableRooms",
			"",
			nil,
			http.MethodGet,
			roomRouter.AvailableRooms,
			http.StatusOK,
		},
	}
)

func TestRoomAccessRouter(t *testing.T) {
	ConfigTest(roomTable, t)
}
