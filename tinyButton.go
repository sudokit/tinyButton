package tinyButton

import (
	"machine"
	"time"
)

type Button struct {
	pin           machine.Pin // pin
	dbTime        int64       // debounce time
	puEnable      bool        // enable pullup resistor
	inverted      bool        // if true, low is pressed, otherwise low is not pressed
	state         bool        // current state (true for pressed and false for not)
	lastState     bool        // last state
	changed       bool        // if the state has changed
	lastTransient bool        // holds variable state during debounce time
	lastDbTime    int64       // last debounce time
	time          int64       // time of the current state
	lastChange    int64       // time of last state change
}

type ToggleButton struct {
	Button      Button
	toggleState bool
}

func NewButton(pin machine.Pin) *Button {
	return &Button{
		pin:           pin,
		dbTime:        25,
		puEnable:      false,
		inverted:      false,
		state:         false,
		lastState:     false,
		lastTransient: false,
		lastDbTime:    0,
		changed:       false,
		time:          0,
		lastChange:    0,
	}
}

func NewToggleButton(pin machine.Pin) *ToggleButton {
	return &ToggleButton{
		Button: Button{
			pin:           pin,
			dbTime:        25,
			puEnable:      false,
			inverted:      true,
			state:         false,
			lastState:     false,
			lastTransient: false,
			lastDbTime:    0,
			changed:       false,
			time:          0,
			lastChange:    0,
		},
		toggleState: false,
	}
}

func (bttn *ToggleButton) Read() bool {
	bttn.Button.Read()
	if bttn.Button.WasPressed() {
		bttn.toggleState = !bttn.toggleState
		bttn.Button.changed = true
	} else {
		bttn.Button.changed = false
	}
	return bttn.toggleState
}

func (bttn *ToggleButton) Configure(dbTime int64, puEnable bool, inverted bool) {
	bttn.Button.Configure(dbTime, puEnable, inverted)
}

func (bttn ToggleButton) Changed() bool { // idk if this is necessary bc button already has a changed function but oh well
	return bttn.Button.changed
}

func (bttn ToggleButton) ToggleState() bool {
	return bttn.toggleState
}

func (bttn *Button) Configure(dbTime int64, puEnable bool, inverted bool) {
	bttn.dbTime = dbTime
	bttn.puEnable = puEnable
	bttn.inverted = inverted

	// idk if to add this to a seperate Begin() function but that seems unnecessary
	if bttn.puEnable {
		bttn.pin.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	} else {
		bttn.pin.Configure(machine.PinConfig{Mode: machine.PinInput})
	}
	bttn.state = bttn.pin.Get()
	if bttn.inverted {
		bttn.state = !bttn.state
	}
	bttn.time = time.Now().UnixMilli()
	bttn.lastState = bttn.state
	bttn.changed = false
	bttn.lastChange = bttn.time
}

func (bttn *Button) Read() bool { // should be called from loop
	ms := time.Now().UnixMilli()
	pinVal := bttn.pin.Get()
	if bttn.inverted {
		pinVal = !pinVal
	}

	if pinVal != bttn.lastTransient {
		bttn.lastDbTime = ms
		bttn.lastTransient = pinVal
		bttn.changed = false
	}
	if (ms - bttn.lastDbTime) > bttn.dbTime {
		bttn.lastState = bttn.state
		bttn.state = pinVal
		bttn.changed = (bttn.state != bttn.lastState)
		if bttn.changed {
			bttn.lastChange = ms
		}
	}
	bttn.time = ms
	return bttn.state
}

func (bttn Button) Changed() bool {
	return bttn.changed
}

func (bttn Button) IsPressed() bool {
	return bttn.state
}

func (bttn Button) IsReleased() bool {
	return !bttn.state
}

func (bttn Button) WasPressed() bool {
	return bttn.state && bttn.changed
}

func (bttn Button) WasReleased() bool {
	return !bttn.state && bttn.changed
}

func (bttn Button) PressedFor(ms int64) bool {
	return bttn.state && bttn.time-bttn.lastChange >= ms
}

func (bttn Button) ReleasedFor(ms int64) bool {
	return !bttn.state && bttn.time-bttn.lastChange >= ms
}

func (bttn Button) LastChange() int64 {
	return bttn.lastChange
}
