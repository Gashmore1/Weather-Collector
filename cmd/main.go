package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Gashmore1/Weather-Collector/pkg/ingest"
)

func main() {
	fmt.Println("Weather Collector starting…")

	var url string

	// Example usage of the fetch function
	if len(os.Args) > 1 {
		url = os.Args[1]
	} else {
		fmt.Println("Usage: weather-collector <URL>")
		return
	}

	ctx := context.Background()
	forecast, err := ingest.FetchForecast(ctx, url)
	if err != nil {
		fmt.Println("Error fetching forecast:", err)
		return
	}
	records, err := ingest.TransformForecast(forecast)
	if err != nil {
		fmt.Println("Error transforming forecast:", err)
		return
	}
	fmt.Printf("Transformed %d records. Sample:\n", len(records))
	for i, r := range records {
		if i >= 5 {
			break
		}
		fmt.Printf("%+v\n", r)
	}
	fmt.Printf("Fetched forecast for latitude %f, longitude %f\n%f\n", forecast.Latitude, forecast.Longitude, forecast.Hourly.Temperature2m[0])

}
