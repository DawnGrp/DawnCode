package main

import (
	"fmt"
	"log"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()

	leds := []*gpio.LedDriver{}

	for i := 1; i <= 40; i++ {
		leds = append(leds, gpio.NewLedDriver(r, fmt.Sprintf("%d", i)))

	}

	work := func() {

		gobot.Every(10*time.Second, func() {
			for _, led := range leds {
				log.Println("led :", led.Name(), led.Pin(), led.State())
				if err := led.Toggle(); err != nil {
					log.Println(led.State())
				}
			}
		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		//[]gobot.Device{leds},
		leds,
		work,
	)

	robot.Start()
}
