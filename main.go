package main

import (
	"github.com/mariuscristian/speedtest/domain"
	"github.com/mariuscristian/speedtest/ports/output/speedtest"
)

func InitMeasurement() domain.SpeedTestUseCase {
	return domain.NewSpeedMeasureUseCase([]domain.SpeedTestClient{
		speedtest.NewOoklaSpeedTestClient(),
		speedtest.NewFastdotcomSpeedTestClient(),
	})
}

func main() {
	speedProbe := InitMeasurement()
	speedProbe.MeasureSpeed(domain.Ookla)
	speedProbe.MeasureSpeed(domain.Netflix)
}
