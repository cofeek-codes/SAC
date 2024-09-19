package mouse

import "github.com/go-vgo/robotgo"

func Click() {
}

func CenterCursor() {
	width, height := robotgo.GetScreenSize()
	robotgo.Move(width/2, height/2)
}
