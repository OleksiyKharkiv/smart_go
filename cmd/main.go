package main

import (
	"fmt"
	"log"
	"smart_go/device"
	"smart_go/power_grid"
)

func main() {
	grid := power_grid.NewPowerGrid(1500)
	kettle := device.NewDevice("Kettle", 1500)
	ac1 := device.NewDevice("AC1", 1000)
	ac2 := device.NewDevice("AC2", 1500)
	unsafeKettle := device.NewDevice("Unsafe Kettle", 2000)

	fmt.Println("\033[32m======= Program Start ========\033[0m")

	err := kettle.TurnOn()
	if err != nil {
		log.Println(err)
	}
	grid.AddDevice(kettle)
	fmt.Printf("%s \n", kettle)

	err = ac1.TurnOn()
	if err != nil {
		log.Println(err)
	}
	grid.AddDevice(ac1)
	fmt.Printf("%s \n", ac1)

	err = ac2.TurnOn()
	if err != nil {
		log.Println(err)
	}
	grid.AddDevice(ac2)
	fmt.Printf("%s \n", ac2)

	unsafeKettle.SetSafeStatus(false)
	err = unsafeKettle.TurnOn()
	if err != nil {
		log.Println(err)
	}
	grid.AddDevice(unsafeKettle)

	fmt.Printf("Total consumption: %d \nMaximal power is: %d\n", grid.TotalConsumption(), grid.MaxPower())

	for grid.TotalConsumption() > grid.MaxPower() {
		disabledDevice, err := grid.AutoDisable()
		if err != nil {
			log.Printf("Error: %v\n", err)
		}
		if disabledDevice == nil {
			break
		}
		fmt.Printf("===============\nPower grid is overloaded!\n Device %s off\n===============\n", disabledDevice.Name())
	}
}
