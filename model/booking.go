package model 


import (
	"strings"
	"math"
)


/*
	creating a struct to how the whole booking 

*/


type Booking struct { 
	pickUp, dropOff 		  string
	destinations              []string
	fare                      Fare
	areaMap               	  map[string]map[string]float64
	transportFare        	  float64
	tip						  float64
}


/*
	creating method to initialize the area map and amount per mile

*/

func (booking *Booking) SetValues() {
	

	booking.SetDestinations([]string{"Choba", "Rumuosi", "Mgbuoba", "Alakahia", "Aluu", "Rumokoro", "Rumuola"})
	areaMap := map[string]map[string]float64{
		"choba": 
		{
			"Lat": 4.8942,
			"Lon": 6.9264,
		},
		"rumuosi": 
		{
			"Lat": 4.8806,
			"Lon": 6.9403,
		},
		"mgbuoba": 
		{
			"Lat": 4.8422,
			"Lon": 6.9692,
		},
		"alakahia": 
		{
			"Lat": 4.8853,
			"Lon": 6.9249,
		},
		"aluu": 
		{
			"Lat": 4.9338,
			"Lon": 6.9435,
		},
		"rumokoro": 
		{
			"Lat": 4.8702,
			"Lon": 6.9883,
		},
		"rumuola": 
		{
			"Lat": 4.8354,
			"Lon": 7.0256,
		},
	}
	booking.SetAreaMap(areaMap)
	booking.fare.SetAmountPerMile(280)
	booking.fare.SetBaseFee(400)
	booking.fare.SetBookingFee(150)
}

/*
	creating method to check the amount entered after the ride is complete 

	return the chage if the amount is greater than the fare

	return minus one if the amount entered is lower than the fare

	return 0 if the amount is equal to the fare

*/


func (book Booking) CheckFareAmountEntered(amount float64) int {
	if fare := book.transportFare; fare < amount {
		//change := amount - fare
		return 1
	} else if fare > amount {
		return -1
	} else 
	{
		return 0
	}
}

/*
	creating method to calculate distance from the pickup to dropoff

*/


func (book Booking)Distance(pickUp, dropOff string, unit ...string) float64 {

	latPickup := book.areaMap[pickUp]["Lat"]
	lngPickup := book.areaMap[pickUp]["Lon"]

	latDropoff := book.areaMap[dropOff]["Lat"]
	lngDropoff := book.areaMap[dropOff]["Lon"]


	const PI float64 = 3.141592653589793
	
	radlat1 := float64(PI * latPickup / 180)
	radlat2 := float64(PI * latDropoff / 180)
	
	theta := float64(lngPickup - lngDropoff)
	radtheta := float64(PI * theta / 180)
	
	dist := math.Sin(radlat1) * math.Sin(radlat2) + math.Cos(radlat1) * math.Cos(radlat2) * math.Cos(radtheta)
	
	if dist > 1 {
		dist = 1
	}
	
	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515
	
	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1.609344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}
	
	totalAmount := book.CalculateFee(dist)

	return totalAmount
}


/*
	creating method to calculate the fee for the total distance from the pickup to dropoff

*/

func (book Booking) CalculateFee(distance float64) float64 {
	amountPerMile := book.fare.GetAmountPerMile()
	baseFee := book.fare.GetBaseFee()
	bookingFee := book.fare.GetBookingFee()

	return (amountPerMile * distance) + baseFee + bookingFee
}



/*
	creating methods to get and set the variable in the booking class

*/
func (book *Booking) SettransportFare(transfare float64) {
	book.transportFare = transfare
}


func (book Booking) GettransportFare() float64 {
	return book.transportFare
}


func (book *Booking) SetDestinations(destination []string) {
	book.destinations = destination
}


func (book Booking) GetDestinations() []string {
	return book.destinations
}


func (book *Booking) SetPickUp(pickUp string) {
	book.pickUp = pickUp
}


func (book Booking) GetPickUp() string {
	return book.pickUp
}


func (book *Booking) SetDropOff(dropOff string) {
	book.dropOff = dropOff
}


func (book Booking) GetDropOff() string {
	return book.dropOff
}


func (book *Booking) SetTip(tip float64) {
	book.tip = tip
}


func (book Booking) GetTip() float64 {
	return book.tip
}


func (book *Booking) SetAreaMap(areaMap map[string]map[string]float64) {
	book.areaMap = areaMap
}


func (book Booking) GetAreaMap() map[string]map[string]float64 {
	return book.areaMap
}


func (book Booking) IsDestinationValid(destination string) bool {
	for i := 0; i < len(book.destinations); i++ {
		if strings.ToLower(book.destinations[i]) == strings.ToLower(destination) {
			return true
			
		}
	}
	return false
}

