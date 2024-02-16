package main

import "fmt"

const ebuKelvin = 373.0

func main() {

	tempKelvin := ebuKelvin
	tempCelsius := tempKelvin - 273.0
	fmt.Printf("A temperatura de ebulição da água em K é %g, e a temperatura de ebulição da água em C é %g.", tempKelvin, tempCelsius)
}
