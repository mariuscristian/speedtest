package domain

import (
	"errors"
	"fmt"
)

type speedMeasureUseCase struct {
	speedTestClients []SpeedTestClient
}

func NewSpeedMeasureUseCase(speedTestClients []SpeedTestClient) SpeedTestUseCase {
	return &speedMeasureUseCase{speedTestClients: speedTestClients}
}

func (s speedMeasureUseCase) MeasureSpeed(method SpeedTestServerType) SpeedTestResult {
	fmt.Printf("starting use case measurement with method %v\n", method)
	for _, client := range s.speedTestClients {
		if client.GetMethod() == method {
			return client.Measure()
		}
	}
	return SpeedTestResult{Err: errors.New("speed test method unknown")}
}
