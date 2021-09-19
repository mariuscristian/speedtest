package tests

import (
	"github.com/mariuscristian/speedtest/domain/measurement"
	"github.com/mariuscristian/speedtest/tests/mocks/domain/measurement"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpeedMeasureUseCase_MeasureSpeed_given_a_measurement_client_when_measurement_method_matches_client_it_succeeds(t *testing.T) {
	mockSpeedTestClient := new(mocks.SpeedTestClient)
	t.Run("success", func(t *testing.T) {
		// given
		const speed = float64(1)
		mockSpeedTestClient.On("GetMethod").Return(measurement.Ookla)
		mockSpeedTestClient.On("Measure").Return(measurement.SpeedTestResult{Download: speed, Upload: speed})
		usecase := measurement.NewSpeedMeasureUseCase([]measurement.SpeedTestClient{mockSpeedTestClient})
		// when
		result := usecase.MeasureSpeed(measurement.Ookla)
		// then
		assert.NoError(t, result.Err)
		assert.Equal(t, result.Download, speed)
		assert.Equal(t, result.Upload, speed)
	})
}

func TestSpeedMeasureUseCase_MeasureSpeed_given_a_measurement_client_when_measurement_method_is_not_matching_with_client_it_fails(t *testing.T) {
	mockSpeedTestClient := new(mocks.SpeedTestClient)
	t.Run("success", func(t *testing.T) {
		// given
		const speed = float64(1)
		mockSpeedTestClient.On("GetMethod").Return(measurement.Netflix)
		mockSpeedTestClient.On("Measure").Return(measurement.SpeedTestResult{Download: speed, Upload: speed})
		usecase := measurement.NewSpeedMeasureUseCase([]measurement.SpeedTestClient{mockSpeedTestClient})
		// when
		result := usecase.MeasureSpeed(measurement.Ookla)
		// then
		assert.Error(t, result.Err)
		assert.Equal(t, result.Download, float64(0))
		assert.Equal(t, result.Upload, float64(0))
	})
}
