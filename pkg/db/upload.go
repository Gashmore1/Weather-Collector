package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Gashmore1/Weather-Collector/pkg/ingest"
)

// UploadRecords inserts hourly weather records into a PostgreSQL database.
// The table is expected to have the following schema:
//
//	CREATE TABLE hourly_records (
//	    id          SERIAL PRIMARY KEY,
//	    record_time TIMESTAMP NOT NULL,
//	    temperature DOUBLE PRECISION,
//	    humidity    INTEGER,
//	    rain        DOUBLE PRECISION,
//	    wind_speed  DOUBLE PRECISION
//	);
//
// Parameters:
//
//	ctx      – context for cancellation and timeouts.
//	db       – an open *sql.DB connection to Postgres.
//	records  – slice of ingest.HourlyRecord to insert.
//
// The function uses a transaction and a prepared statement for safety.
// It returns an error if any part of the operation fails.
func UploadRecords(ctx context.Context, db *sql.DB, records []ingest.HourlyRecord) error {
	if len(records) == 0 {
		return nil
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	// Rollback if anything goes wrong.
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("transaction rollback failed: %v", rbErr)
			}
		}
	}()

	stmt, err := tx.PrepareContext(ctx, `INSERT INTO hourly_records (record_time, temperature, humidity, rain, wind_speed) VALUES ($1, $2, $3, $4, $5)`)
	if err != nil {
		return fmt.Errorf("prepare statement: %w", err)
	}
	defer stmt.Close()

	for _, r := range records {
		// Parse the timestamp string into time.Time. Assume RFC3339 format.
		if _, err = stmt.ExecContext(ctx, r.Time, r.Temperature, r.Humidity, r.Rain, r.WindSpeed); err != nil {
			return fmt.Errorf("exec insert: %w", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}
	return nil
}
