package examples

import (
	"fmt"
	"gobot.io/x/gobot"
	_ "gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func stepperManager(stepper *gpio.StepperDriver, fa *firmata.Adaptor) {
	fmt.Println("Init stepper")
	for {
		hit, err := fa.DigitalRead("4")
		if err != nil {
			panic("Could not read pin 4")
		}

		if hit == 0 {
			stepper.Move(-300)
			break
		} else {
			stepper.Move(10)
		}
	}
	for i := 0; i < 100; i++ {
		stepper.Move(-50)
	}
	for i := 0; i < 100; i++ {
		stepper.Move(50)
	}
}

func StepByStep() *gobot.Robot {

	firmataAdapter := firmata.NewAdaptor("/dev/ttyACM0")

	stepper := gpio.NewStepperDriver(
		firmataAdapter,
		[4]string{"8", "9", "10", "11"},
		gpio.StepperModes.DualPhaseStepping,
		2048)
	work := func() {
		stepperManager(stepper, firmataAdapter)
	}

	robot := gobot.NewRobot("Stepper",
		[]gobot.Connection{firmataAdapter},
		[]gobot.Device{stepper},
		work,
	)
	return robot
}
