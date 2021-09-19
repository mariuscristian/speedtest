package tests

import (
	"github.com/mariuscristian/speedtest/config"
	"github.com/mariuscristian/speedtest/domain/measurement"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpeedWithOokla(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		result := config.InitMeasurements().MeasureSpeed(measurement.Ookla)
		assert.NoError(t, result.Err)
		assert.NotZero(t, result.Download)
		assert.NotZero(t, result.Upload)
	})
}

func TestSpeedWithNetflix(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		result := config.InitMeasurements().MeasureSpeed(measurement.Netflix)
		assert.NoError(t, result.Err)
		assert.NotZero(t, result.Download)
		assert.NotZero(t, result.Upload)
	})
}
