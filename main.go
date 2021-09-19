package main

import (
	"github.com/mariuscristian/speedtest/measurement"
	"github.com/mariuscristian/speedtest/ports/input"
)

func main() {
	speedProbe := init.Measurements()
	speedProbe.MeasureSpeed(measurement.Ookla)
	speedProbe.MeasureSpeed(measurement.Netflix)
}
