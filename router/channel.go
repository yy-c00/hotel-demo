package router

import (
	"github.com/olahol/melody"
)

func connect(session *melody.Session) {
	header := session.Request.Header

	user, room := header.Get("User-Data"), header.Get("Room-Data")

	session.Set("user", user)
	session.Set("room", room)
}