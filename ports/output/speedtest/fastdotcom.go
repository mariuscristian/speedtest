package speedtest

import (
	"context"
	"errors"
	"github.com/mariuscristian/speedtest/domain"
	"go.jonnrb.io/speedtest/fastdotcom"
	"go.jonnrb.io/speedtest/units"
	"log"
	"time"
)

const (
	urlCount int     = 5
	Kbps     float64 = 1000
	Mbps             = 1000 * Kbps
)

type fastdotcomSpeedTestClient struct {
	downloadSpeed float64
	uploadSpeed   float64
}

func NewFastdotcomSpeedTestClient() domain.SpeedTestClient {
	log.Println("[netflix] creating client")
	return &fastdotcomSpeedTestClient{}
}

// Measure - performs the speed measurement
func (fc fastdotcomSpeedTestClient) Measure() domain.SpeedTestResult {
	log.Println("[netflix] starting measurement")
	var client fastdotcom.Client

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	m, err := fastdotcom.GetManifest(ctx, urlCount)
	if err != nil {
		return domain.SpeedTestResult{Err: errors.New("fast.com speed test config failed")}
	}

	err = fc.testDownload(m, &client)
	if err != nil {
		return domain.SpeedTestResult{Err: errors.New("fast.com speed test download failed")}
	}

	err = fc.testUpload(m, &client)
	if err != nil {
		return domain.SpeedTestResult{Err: errors.New("fast.com speed test upload failed")}
	}

	return domain.SpeedTestResult{Download: fc.downloadSpeed, Upload: fc.uploadSpeed}
}

// GetMethod - checks the api used for the speed measurement
func (fc fastdotcomSpeedTestClient) GetMethod() domain.SpeedTestServerType {
	return domain.Netflix
}

func (fc *fastdotcomSpeedTestClient) testDownload(m *fastdotcom.Manifest, client *fastdotcom.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream := make(chan units.BytesPerSecond)
	go func() {
		for range stream {
		}
	}()

	speed, err := m.ProbeDownloadSpeed(ctx, client, stream)
	if err != nil {
		return err
	}

	fc.downloadSpeed = float64(speed.BitsPerSecond()) / Mbps

	return nil
}

func (fc *fastdotcomSpeedTestClient) testUpload(m *fastdotcom.Manifest, client *fastdotcom.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream := make(chan units.BytesPerSecond)
	go func() {
		for range stream {
		}
	}()

	speed, err := m.ProbeUploadSpeed(ctx, client, stream)
	if err != nil {
		return err
	}

	fc.uploadSpeed = float64(speed.BitsPerSecond()) / Mbps

	return nil
}
