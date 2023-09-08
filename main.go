package main

import (
	"blinky/button"
	"machine"
)

// func usToMs(us float32) int16 {
// 	return int16(us * 1000)
// }

// var (
// 	pwm = machine.Timer1
// 	pin = machine.D9
// )

func main() {
	// s, err := servo.New(pwm, pin)
	// if err != nil {
	// 	for {
	// 		println("could not configure servo")
	// 		time.Sleep(time.Second)
	// 	}
	// 	return
	// }

	bttn_pin := machine.D5
	button := button.NewButton(bttn_pin)
	button.Configure(25, true, true)
	// button.Begin()

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// println("setting to 45°")
	// s.SetMicroseconds(usToMs(1))
	// time.Sleep(3 * time.Second)

	// println("setting to 90°")
	// s.SetMicroseconds(usToMs(1.5))
	// time.Sleep(3 * time.Second)

	// println("setting to 135")
	// s.SetMicroseconds(usToMs(2))
	// time.Sleep(3 * time.Second)

	// println("setting to 180")
	// s.SetMicroseconds(usToMs(2.5))
	// time.Sleep(3 * time.Second)

	for {
		// bttn_state := button.Read()
		button.Read()
		// if button.IsPressed() {
		// 	led.High()
		// } else {
		// 	led.Low()
		// }
		// if button.Changed() {
		// 	led.Set(bttn_state)
		// 	println("Button state is ", bttn_state)
		// }
		println(button.Changed())
		// button.Changed()
	}
}
