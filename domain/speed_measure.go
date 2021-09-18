package domain

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
	MeasureSpeed(method SpeedTestServerType) SpeedTestResult
}

type SpeedTestClient interface {
	Measure() SpeedTestResult
	GetMethod() SpeedTestServerType
}
