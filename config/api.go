package config

import (
	"github.com/mariuscristian/speedtest/measurement"
	"github.com/mariuscristian/speedtest/ports/output/speedtest"
)

func InitMeasurements() measurement.SpeedTestUseCase {
	return measurement.NewSpeedMeasureUseCase([]measurement.SpeedTestClient{
		speedtest.NewOoklaSpeedTestClient(),
		speedtest.NewFastdotcomSpeedTestClient(),
	})
}
