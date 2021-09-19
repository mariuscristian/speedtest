package domain

import (
	"errors"
	"log"
)

type speedMeasureUseCase struct {
	speedTestClients []SpeedTestClient
}

func NewSpeedMeasureUseCase(speedTestClients []SpeedTestClient) SpeedTestUseCase {
	return &speedMeasureUseCase{speedTestClients: speedTestClients}
}

// MeasureSpeed searches the chosen client and delegates the speed measurement
func (s speedMeasureUseCase) MeasureSpeed(method SpeedTestServerType) SpeedTestResult {
	log.Printf("[UC measure speed] starting measurement with method %v\n", method)
	for _, client := range s.speedTestClients {
		if client.GetMethod() == method {
			result := client.Measure()
			log.Printf("[UC measure speed] results - Download: %f, Upload: %f, Error: %v\n", result.Download, result.Upload, result.Err)
			return client.Measure()
		}
	}
	return SpeedTestResult{Err: errors.New("speed test method unknown")}
}
