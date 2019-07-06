package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()

	rec := gpio.NewDirectPinDriver(r, "8")
	ctl := gpio.NewDirectPinDriver(r, "10")

	work := func() {
		gobot.Every(1*time.Second, func() {

			ctl.DigitalWrite(1)

			time.Sleep(20 * time.Microsecond)

			ctl.DigitalWrite(0)

			toggle := false
			var recTime int64
			var long int64
			for {
				i, err := rec.DigitalRead()
				if err != nil {
					fmt.Println(err.Error())
				}

				if i == 1 {
					if toggle == false {
						recTime = time.Now().UnixNano()
						toggle = true
						fmt.Println("record")
					}

				} else {
					if toggle == true {
						long = time.Now().UnixNano() - recTime
						toggle = false
						fmt.Println("end record:", long, (long*340000)/1000000000, "mm")
						break
					}

				}

			}

		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		[]gobot.Device{rec, ctl},

		work,
	)

	robot.Start()
}
