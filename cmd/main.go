package main

import (
	"fmt"
	"log"
	"smart_go/device"
)

func main() {
	kettle := device.NewDevice("Kettle", 1500)
	ac1 := device.NewDevice("AC1", 1000)
	ac2 := device.NewDevice("AC2", 1500)

	fmt.Println("======= Program Start  ========")
	log.Println("Begin...")

	err := kettle.TurnOn()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("device %s is ON", kettle)

	err = ac1.TurnOn()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("device %s is ON", ac1)

	err = ac2.TurnOn()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("device %s is ON", ac2)
}
