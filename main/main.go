package main

import (
	"cabservice/model"
	"fmt"
	"math"
	"os"
	"strings"
	"bufio"
	"time"
)



func main() {

	initializeValues()

}

func initializeValues(){
	// creating an instance of the booking struct created in the booking.go file in models package
		ride:= model.Booking{}

	// initializing the route comuputed and the price per mile and other stuff
		ride.SetValues()


		rideDestination := ride.GetDestinations()
		fmt.Println("Welcome, these are the destinations we commute: ")
		fmt.Println(rideDestination)


		pickup := getPickUpLocation()
		ride.SetPickUp(pickup)
		isPupValid := ride.IsDestinationValid(pickup)
		isDestinationValid(isPupValid, 1)

		dropOff := getDropOffLocation()
		ride.SetDropOff(dropOff)
		isDolValid := ride.IsDestinationValid(dropOff)
		isDestinationValid(isDolValid, 2)


		transportFare := getTransportAmount(ride,pickup,dropOff)
		ride.SettransportFare(transportFare)

		fmt.Printf("Total cost of ₦%.2f to go from %v to %v\n", transportFare, pickup, dropOff)
		time.Sleep(10 * time.Second)
		value, amount := collectMoney(ride)

		change := calculateChange(value,amount,transportFare)
		fmt.Printf("Here's your change of ₦%.2f\n", change)

		fmt.Println("Baba find your guy something...")
		_, info := collectTip(transportFare)
		fmt.Println(info)

}

func collectTip(fare float64) (float64, string) {
	var tips float64
	var info string
	a := 1
	fmt.Println("Enter money you won find your guy:")

	for a == 1 {
		_, err := fmt.Scan(&tips)
		if err != nil {
			fmt.Println("ah you no won find me something sho, abeg shake body ni:")
			a = 1
			continue
		}

		if tips <= 0 {
			info = "Ah this one is stingy oo, local man can not can."
			a = -1
			tips = 0
		} else if tips > 0 && tips <= fare {
			info = "Gracias Mucho!!!!!"
			a = -1
		} else {
			info = "Gracias Mucho!!!!!"
			a = -1
		}
	}
	return (math.Floor(tips*100) / 100), info
}


func getPickUpLocation() string{
		pickUpLocation := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your pickup location:\n")
		pul, _ := pickUpLocation.ReadString('\n')

		pickUpP := strings.ToLower(strings.Trim(pul, " \r\n"))

		return pickUpP
}

func getDropOffLocation() string {
		dropOffLocation := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your drop off Location:")
		dol, _ := dropOffLocation.ReadString('\n')

		dropOffP := strings.ToLower(strings.Trim(dol, " \r\n"))

		return dropOffP
}

func getTransportAmount(ride model.Booking,pickup,dropOff string) float64{
		transportFare := ride.Distance(pickup, dropOff,"M")
		transportFare = (math.Floor(transportFare*100) / 100)

		return transportFare
}

func collectMoney(ride model.Booking) (int, float64) {
	a := -1
	var amount float64
	count := 0
	fmt.Println("Enter amount to pay:")

	for a < 0 {
		_, err := fmt.Scan(&amount)
		if err != nil {
			fmt.Println("You entered an invalid amount...")
			fmt.Println("Re-enter amount to pay:")
			continue
		}

		a = ride.CheckFareAmountEntered(amount)

		if a == -1 {
			fmt.Println("No money  here...")
			fmt.Println("Re-enter amount to pay:")
		}

		count++

		if count%5 == 0 {
			fmt.Println("You'll be reported to the police if you keep on trying to pay less than actual amount\n")
			fmt.Println("Re-enter amount to pay:")
			continue
		}
	}
	return a, (math.Floor(amount*100) / 100)
}

func calculateChange(value  int, amount, transportFare float64) float64 {
	var change float64
	if value == 1 {
		change = amount - transportFare
	}
	return change
}

func isDestinationValid(val bool, spot int) {
	if val == false && spot == 1 {
		fmt.Println("Invalid Pickup Destination...\n")
		fmt.Println("Bye laters...")
		os.Exit(1)
	} else if val == false && spot == 2 {
		fmt.Println("Invalid Dropoff Destination\n...")
		fmt.Println("Bye laters...")
		os.Exit(1)
	}
}
