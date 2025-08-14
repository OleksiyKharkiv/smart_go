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
	unsafeKettle := device.NewDevice("Unsafe Kettle", 2000)

	fmt.Println("\033[32m======= Program Start ========\033[0m")

	err := kettle.TurnOn()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%s \n", kettle)

	err = ac1.TurnOn()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%s \n", ac1)

	err = ac2.TurnOn()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%s \n", ac2)

	unsafeKettle.SetSafeStatus(false)
	err = unsafeKettle.TurnOn()
	if err != nil {
		log.Println(err)
	}
}
