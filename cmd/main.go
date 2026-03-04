package main

import (
	"context"
	"fmt"
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/Gashmore1/Weather-Collector/pkg/db"
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


	username := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	database_hostname := os.Getenv("DATABASE_URL")
	database_port := os.Getenv("DATABASE_PORT")
	database_name := os.Getenv("DATABASE_NAME")

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, database_hostname, database_port, database_name)

	// Upload to Postgres
	dbConn, err := sql.Open("pgx", connectionString)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer dbConn.Close()

	if err := db.UploadRecords(ctx, dbConn, records); err != nil {
		fmt.Println("Error uploading records:", err)
		return
	}
	fmt.Println("Uploaded records successfully.")
	fmt.Printf("Fetched forecast for latitude %f, longitude %f\n%f\n", forecast.Latitude, forecast.Longitude, forecast.Hourly.Temperature2m[0])
}
