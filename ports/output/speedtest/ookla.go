package speedtest

import (
	"errors"
	"fmt"
	"github.com/mariuscristian/speedtest/domain"
	"github.com/showwin/speedtest-go/speedtest"
)

type ooklaSpeedTestClient struct {
}

func NewOoklaSpeedTestClient() domain.SpeedTestClient {
	fmt.Println("creating ookla client")
	return &ooklaSpeedTestClient{}
}

func (o ooklaSpeedTestClient) Measure() domain.SpeedTestResult {
	fmt.Println("starting ookla measurement")
	user, _ := speedtest.FetchUserInfo()
	serverList, _ := speedtest.FetchServerList(user)
	targets, _ := serverList.FindServer([]int{})

	for _, s := range targets {
		fmt.Printf("ookla measurement for host %v\n", s.Host)
		s.PingTest()
		s.DownloadTest(false)
		s.UploadTest(false)
		fmt.Printf("Latency: %s, Download: %f, Upload: %f\n", s.Latency, s.DLSpeed, s.ULSpeed)
		return domain.SpeedTestResult{Download: s.DLSpeed, Upload: s.ULSpeed}
	}
	return domain.SpeedTestResult{Err: errors.New("ookla speed test failed")}
}

func (o ooklaSpeedTestClient) GetMethod() domain.SpeedTestServerType {
	return domain.Ookla
}
