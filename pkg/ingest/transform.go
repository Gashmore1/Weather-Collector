package ingest

import (
	"fmt"
	"time"

	"github.com/Gashmore1/Weather-Collector/pkg/model"
)

// HourlyRecord represents a single row of hourly weather data.
type HourlyRecord struct {
	Time        time.Time
	Temperature float64
	Humidity    int
	Rain        float64
	WindSpeed   float64
}

var LAYOUT string = "2006-01-02T15:04"

// TransformForecast pivots the column-oriented Forecast data into a slice of HourlyRecord rows.
// It returns an error if the slice lengths are inconsistent.
func TransformForecast(f *model.Forecast) ([]HourlyRecord, error) {
	n := len(f.Hourly.Time)
	if n == 0 {
		return nil, nil
	}
	// Validate that all slices are the same length.
	if len(f.Hourly.Temperature2m) != n || len(f.Hourly.RelativeHumidity2m) != n ||
		len(f.Hourly.Rain) != n || len(f.Hourly.WindSpeed10m) != n {
		return nil, fmt.Errorf("forecast data slices have mismatched lengths")
	}
	out := make([]HourlyRecord, n)
	for i := 0; i < n; i++ {
		if date, err := time.Parse(LAYOUT, f.Hourly.Time[i]); err != nil {
			fmt.Printf("Failed to convert time data from string to time: %s", err.Error())
		} else {
			out[i] = HourlyRecord{
				Time:        date,
				Temperature: f.Hourly.Temperature2m[i],
				Humidity:    f.Hourly.RelativeHumidity2m[i],
				Rain:        f.Hourly.Rain[i],
				WindSpeed:   f.Hourly.WindSpeed10m[i],
			}
		}
	}
	return out, nil
}
