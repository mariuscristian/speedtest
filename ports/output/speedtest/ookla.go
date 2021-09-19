package speedtest

import (
	"errors"
	"github.com/mariuscristian/speedtest/measurement"
	"github.com/showwin/speedtest-go/speedtest"
	"log"
)

type ooklaSpeedTestClient struct {
}

func NewOoklaSpeedTestClient() measurement.SpeedTestClient {
	log.Println("[ookla] creating client")
	return &ooklaSpeedTestClient{}
}

// Measure - performs the speed measurement
func (c ooklaSpeedTestClient) Measure() measurement.SpeedTestResult {
	log.Println("[ookla] starting measurement")
	user, _ := speedtest.FetchUserInfo()
	serverList, _ := speedtest.FetchServerList(user)
	targets, _ := serverList.FindServer([]int{})

	for _, s := range targets {
		log.Printf("[ookla] measurement for host %v\n", s.Host)
		err := s.PingTest()
		if err != nil {
			return measurement.SpeedTestResult{Err: errors.New("ookla ping test failed")}
		}
		err = s.DownloadTest(false)
		if err != nil {
			return measurement.SpeedTestResult{Err: errors.New("ookla download speed test failed")}
		}
		err = s.UploadTest(false)
		if err != nil {
			return measurement.SpeedTestResult{Err: errors.New("ookla upload speed test failed")}
		}

		return measurement.SpeedTestResult{Download: s.DLSpeed, Upload: s.ULSpeed}
	}
	return measurement.SpeedTestResult{Err: errors.New("ookla speed test failed")}
}

// GetMethod - checks the api used for the speed measurement
func (c ooklaSpeedTestClient) GetMethod() measurement.SpeedTestServerType {
	return measurement.Ookla
}
