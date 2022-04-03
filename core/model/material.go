package model

// Material ...
type Material struct {
	CompositeKey      string
	Location          string
	Pile              string
	Position          string
	OptimumCapacity   float64
	MaxCapacity       float64
	AreaName          string
	AreaCategory      string
	Contractor        string
	Key               string
	Date              string
	Day               int32
	Month             int32
	Year              int32
	WorkDay           int32
	Aging             int32
	Type              string
	Grade             string
	Actual            float64
	Target            float64
	Troughput         float64
	PreviousActual    float64
	PreviousTroughput float64
	TotalActual       float64
	Throughput        float64
	Value             float64
	Temperature       float64
	Distance          float64
	TargetThroughput  float64
}

type PerformanceRawMaterial struct {
	Previous map[string]*Material
	Current  map[string]*Material
}

type RawMaterial struct {
	Previous []*Material
	Current  []*Material
}

type PerformanceMaterial struct {
	WasteValue             map[string]*Material
	ProductionValue        map[string]*Material
	WasteAdjValue          map[string]*Material
	ProductionAdjValue     map[string]*Material
	WastePrevValue         map[string]*Material
	ProductionPrevValue    map[string]*Material
	WastePrevAdjValue      map[string]*Material
	ProductionPrevAdjValue map[string]*Material
	WasteTargetValue       map[string]*Material
	ProductionTargetValue  map[string]*Material
	Distance               float64
	ProductBrand           map[string][]*Material
}

// Summary ...
type Summary struct {
	Year        int32
	Month       int32
	Value       float64
	Type        int32
	Grade       string
	Key         string
	AvgDelivery float64
	Count       int32
}
