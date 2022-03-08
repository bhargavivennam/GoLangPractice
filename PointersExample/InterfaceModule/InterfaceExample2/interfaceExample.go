package main

import "fmt"

type R interface {
	getMenu() string
	getSpecials() string
}

type Restaurant1 struct {
	itemName     string
	liveTracking string
	arrivalTime  string
	specialItem  string
}

type Restaurant2 struct {
	itemName    string
	arrivalTime string
	onlineOrder string
	specialItem string
}

func main() {

	X := Restaurant1{
		itemName:     "chicken biryani",
		liveTracking: "Your order is 2 miles away",
		arrivalTime:  "Arrival time is 5:00PM",
		specialItem:  "Mutton Biryani",
	}

	y := Restaurant2{
		itemName:    "idly and chutney",
		arrivalTime: "Your order is handed over to the uber agent",
		onlineOrder: "Your order has been placed",
		specialItem: "Chilly Chicken",
	}

	fmt.Print("Choose one of the following:\n1.Restaurant 1\n2.Restaurant 2\n")
	var op int
	fmt.Scan(&op)
	switch op {
	case 1:
		fmt.Println("the menu is \n", Menu(X))
		fmt.Println("Today's special item is:", getSpecials(X))

	case 2:
		fmt.Println("the menu is ", Menu(y))
		fmt.Println("Today's special item is:", getSpecials(y))
	}
}

func Menu(r R) string {
	return r.getMenu()
}

func (b Restaurant1) getMenu() string {
	return b.itemName + "\n" + b.arrivalTime + "\n" + b.liveTracking
}

func (c Restaurant2) getMenu() string {
	return c.itemName + "\n" + c.arrivalTime + "\n" + c.onlineOrder
}
func getSpecials(s R) string {
	return s.getSpecials()
}
func (d Restaurant1) getSpecials() string {
	return d.specialItem
}
func (e Restaurant2) getSpecials() string {
	return e.specialItem
}
