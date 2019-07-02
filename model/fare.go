package model

type Fare struct {
	amountPerMile, bookingFee, baseFee float64
}

func (f *Fare) SetAmountPerMile(costPerMile float64) {
	f.amountPerMile = costPerMile
}

func (f Fare) GetAmountPerMile() float64 {
	return f.amountPerMile
}

func (f *Fare) SetBookingFee(bookingFee float64) {
	f.bookingFee = bookingFee
}

func (f Fare) GetBookingFee() float64 {
	return f.bookingFee
}

func (f *Fare) SetBaseFee(baseFee float64) {
	f.baseFee = baseFee
}

func (f Fare) GetBaseFee() float64 {
	return f.baseFee
}