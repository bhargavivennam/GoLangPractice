package main

import "fmt"

// type geometry interface {
// 	Rectangle() int
// 	Triangle() float64
// 	Square() int
// }

// type Parameters struct {
// 	length  float64
// 	breadth float64
// }

// func (p *Parameters) Rectangle() float64 {

// 	area := p.length * p.breadth
// 	return float64(area)
// }

// func (p *Parameters) Triangle() float64 {

// 	area1 := float64(0.5) * p.length * p.breadth
// 	return float64(area1)
// }

// func (p *Parameters) Square() float64 {
// 	area2 := p.length * p.length
// 	return float64(area2)
// }

// func main() {
// 	X := Parameters{
// 		length:  20,
// 		breadth: 20,
// 	}

// 	fmt.Println(X.Rectangle())
// 	fmt.Println(X.Triangle())
// 	fmt.Println(X.Square())
// }

//implement speed of a car , bus , train and bike vehicles

type vehicles interface {
	CarDemo()
	BusDemo()
	TrainDemo()
	BikeDemo()
}

// type speedDemo struct {
// 	speedofCar   int
// 	speedofBus   int
// 	speedofTrain int
// 	speedofBike  int
// }

type Car struct {
	name        string
	color       string
	tyres       int
	speed       int
	destination string
	count       int
}

type Bus struct {
	color       string
	tyres       int
	speed       int
	destination string
}

type Train struct {
	color       string
	speed       int
	destination string
}

type Bike struct {
	name        string
	color       string
	destination string
}

// func (c *Car) CarDemo() {
// 	res := Car{
// 		name:        "Ferrari",
// 		color:       "red",
// 		tyres:       4,
// 		speed:       200,
// 		destination: "texas",
// 		count:       2,
// 	}
// 	fmt.Println(res)
// }

// func (b *Bus) BusDemo() {
// 	res1 := Bus{
// 		color:       "Black",
// 		tyres:       8,
// 		speed:       300,
// 		destination: "New Jersey",
// 	}
// 	fmt.Println("Details of bus:", res1)
// }

// func (t *Train) TrainDemo() {
// 	res2 := Train{
// 		color:       "Blue",
// 		speed:       400,
// 		destination: "Missouri",
// 	}
// 	fmt.Println("Details of Train:", res2)
// }

// func (bi *Bike) BikeDemo() {
// 	bib := &Bike{
// 		name:        " Enfield",
// 		color:       "Green",
// 		destination: "Irvig",
// 	}
// 	fmt.Println("Details of Bike:", bib)
// }

func main() {
	//var v vehicles
	b := &Bike{
		name:        "Royal Enfield",
		color:       "Green",
		destination: "Irving",
	}
	//v = b

	t := &Train{
		color:       "Blue",
		speed:       400,
		destination: "Missouri",
	}
	//v = t

	bu := &Bus{
		color:       "Black",
		tyres:       8,
		speed:       300,
		destination: "New Jersey",
	}
	//v = bu

	c := &Car{
		name:        "Ferrari",
		color:       "red",
		tyres:       4,
		speed:       200,
		destination: "texas",
		count:       2,
	}

	//v = c

	fmt.Println(b, t, bu, c)
	// v.BikeDemo()
	//v = c
	//v.

	// v.BusDemo()
	// v.CarDemo()
	// v.TrainDemo()
}

// func (s speedDemo) Car() {
// 	fmt.Println("speed of car: ", s.speedofCar)
// }

// func (s speedDemo) Bus() {
// 	fmt.Println("speed of bus: ", s.speedofBus)

// }

// func (s speedDemo) Train() {
// 	fmt.Println("speed of train: ", s.speedofTrain)

// }

// func (s speedDemo) Bike() {
// 	fmt.Println("speed of bike: ", s.speedofBike)

// }

// func main() {
// 	result := speedDemo{
// 		speedofCar:   120,
// 		speedofBus:   200,
// 		speedofTrain: 300,
// 		speedofBike:  60,
// 	}
// 	fmt.Println(result)
// }
