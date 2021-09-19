# speedtest
This module can be used to test download and upload speeds using
- Ookla's https://www.speedtest.net/ and
- Netflix's https://fast.com/

## Go API

```
go get github.com/mariuscristian/speedtest
```

## API Usage

Initialize the speed test use cases
```go
package main

import (
	"github.com/mariuscristian/speedtest"
)
func main() {
	speedProbe := InitMeasurement()
	speedProbe.MeasureSpeed(domain.Ookla)
	speedProbe.MeasureSpeed(domain.Netflix)
}
```

