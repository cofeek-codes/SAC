package clicker

import (
	"log"
	"time"
)

type Clicker struct {
	ToggleKey string
	ExitKey   string
	Delay     int
	IsActive  bool
}

func NewClicker() *Clicker {
	return &Clicker{
		ToggleKey: "a",
		ExitKey:   "b",
		Delay:     1000, // delay is represented in milliseconds
		IsActive:  false,
	}
}

func (c *Clicker) Launch() {
	// @TODO: add system level notification
	log.Println("Clicker launched in background")
	go c.run()
}

func (c *Clicker) Activate() {
	c.IsActive = true
	log.Println("Clicker activated")
}

func (c *Clicker) Deactivate() {
	c.IsActive = false
	log.Println("Clicker deactivated")
}

func (c *Clicker) Exit() {
	if c.IsActive {
		c.Deactivate()
	}
}

func (c *Clicker) run() {
	for c.IsActive {
		log.Println("Click")
		time.Sleep(time.Duration(c.Delay) * time.Millisecond)
	}
}
