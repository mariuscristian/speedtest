# speedtest
This module can be used to test download and upload speeds using
- Ookla's https://www.speedtest.net/ and
- Netflix's https://fast.com/

## Go API

```
go get github.com/mariuscristian/speedtest
```

## API Usage

Initialize the speed test use cases and then run the test by the method type

```go
package main

import (
	"github.com/mariuscristian/speedtest/config"
	"github.com/mariuscristian/speedtest/domain/measurement"
)

func main() {
	speedProbe := config.InitMeasurements()
	speedProbe.MeasureSpeed(measurement.Ookla)
	speedProbe.MeasureSpeed(measurement.Netflix)
}
```

