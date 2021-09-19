package config

import (
	measurement2 "github.com/mariuscristian/speedtest/domain/measurement"
	"github.com/mariuscristian/speedtest/ports/output/speedtest"
)

func InitMeasurements() measurement2.SpeedTestUseCase {
	return measurement2.NewSpeedMeasureUseCase([]measurement2.SpeedTestClient{
		speedtest.NewOoklaSpeedTestClient(),
		speedtest.NewFastdotcomSpeedTestClient(),
	})
}
