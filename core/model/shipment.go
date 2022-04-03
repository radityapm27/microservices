package model

// DeliveryProduct ...
type DeliveryProduct struct {
	Key        string
	Day        int32
	Month      int32
	Year       int32
	Brand      string
	Actual     float64
	Target     float64
	PrevActual float64
}

// DeliveryRatio ...
type DeliveryRatio struct {
	Brand string
	Month int32
	Year  int32
	Ratio float64
}

// TranshipmentSchedule ...
type TranshipmentSchedule struct {
	Year   int32
	Dem    float64
	Des    float64
	Month  int32
	Day    int32
	Amount float64
	Buyer  string
	Vessel string
	Laycan string
	Eta    string
	Brand  string
	Key    string
}

// SalesProduct ...
type SalesProduct struct {
	Grade        string
	AvgSellPrice float64
	AvgDelivery  float64
	Actual       float64
	Revenue      float64
}

// MotherVessel ...
type MotherVessel struct {
	Name        string
	Customer    string
	Target      float64
	Actual      float64
	QueueNumber int32
	Brand       string
	RiskStatus  string
	Type        string
}
