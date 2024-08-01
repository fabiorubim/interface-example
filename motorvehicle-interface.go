package main

import "fmt"

type Motorvehicle interface {
	Mileage() float64
	AverageSpeed() float64
	// StartEngine()
}

type BMW struct {
	distance     float64
	fuel         float64
	averagespeed float64
}

func (b BMW) Mileage() float64 {
	return b.distance / b.fuel
}

func (b BMW) AverageSpeed() float64 {
	return b.averagespeed
}

func ExecuteMotorvehicleInterface() {
	var m Motorvehicle
	m = BMW{distance: 63.0, fuel: 14.4, averagespeed: 46.9}
	q := BMW{distance: 98.0, fuel: 26.3, averagespeed: 52.9}
	fmt.Printf("type of m is %T\n", m)
	fmt.Printf("value of m is %v\n", m)
	fmt.Println("Mileage of 1st ride m", m.Mileage())
	fmt.Println("Mileage of 2st ride 1", q.Mileage())
}
