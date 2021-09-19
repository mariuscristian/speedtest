package main

import (
	"github.com/mariuscristian/speedtest/config"
	"github.com/mariuscristian/speedtest/measurement"
)

func main() {
	speedProbe := config.InitMeasurements()
	speedProbe.MeasureSpeed(measurement.Ookla)
	speedProbe.MeasureSpeed(measurement.Netflix)
}
