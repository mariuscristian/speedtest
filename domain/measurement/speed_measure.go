package measurement

type SpeedTestServerType int

const (
	Ookla SpeedTestServerType = iota
	Netflix
)

type SpeedTestResult struct {
	Download float64
	Upload   float64
	Err      error
}

type SpeedTestUseCase interface {
	// MeasureSpeed searches the chosen client and delegates the speed measurement
	MeasureSpeed(method SpeedTestServerType) SpeedTestResult
}

type SpeedTestClient interface {
	// Measure - performs the speed measurement
	Measure() SpeedTestResult
	// GetMethod - checks the api used for the speed measurement
	GetMethod() SpeedTestServerType
}
