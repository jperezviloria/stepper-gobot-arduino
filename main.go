package main

import (
	"fmt"
	"github.com/jperezviloria/leds/examples"
	"gobot.io/x/gobot"
)

func main() {
	//examples.Hello()
	//examples.Led()
	//examples.Led2()
	//examples.StepByStep()

	master := gobot.NewMaster()
	master.AddRobot(examples.StepByStep())
	fmt.Println("Starting master")
	master.Start()
}
