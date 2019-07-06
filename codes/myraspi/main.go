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
	ledC := gpio.NewLedDriver(r, "8")
	ledR := gpio.NewLedDriver(r, "10")

	work := func() {
		gobot.Every(2*time.Second, func() {

			go func() {
				for {

					if ledR.State() {
						fmt.Println("receiveing")
					}
				}
			}()

			fmt.Println("send start")
			err := ledC.On()
			fmt.Println("ledc on", ledC.State(), ledR.State())
			if err != nil {
				fmt.Println(err.Error())
			}
			time.Sleep(10 * time.Millisecond)
			err = ledC.Off()
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println("send over")

			fmt.Println("ledc off", ledC.State(), ledR.State())

		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		[]gobot.Device{ledC, ledR},
		work,
	)

	robot.Start()
}
