package main

import (
	"github.com/mariuscristian/speedtest/domain"
	"github.com/mariuscristian/speedtest/ports/output/speedtest"
)

func main() {
	ookla := speedtest.NewOoklaSpeedTestClient()
	speed := domain.NewSpeedMeasureUseCase([]domain.SpeedTestClient{ookla})
	speed.MeasureSpeed(domain.Ookla)
}
