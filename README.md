# speedtest
This module can be used to test download and upload speeds using
- Ookla's https://www.speedtest.net/ and
- Netflix's https://fast.com/

## Go API
To use the library in your go application, get the package through
```
go get -d github.com/mariuscristian/speedtest
```

## API Usage
Initialize the speed test use cases and then run the test using the 2 available measurement types

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

