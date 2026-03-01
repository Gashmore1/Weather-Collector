package ingest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Gashmore1/Weather-Collector/pkg/model"
)

// FetchForecast retrieves the JSON from the given URL and unmarshals it into a model.Forecast.
// It returns an error if the HTTP request fails, the status is non‑200, or the body cannot be decoded.
func FetchForecast(ctx time.Context, url string) (*model.Forecast, error) {
	// Create an HTTP request with the provided context so callers can cancel.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	// Perform the request.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Read a bit of the body for diagnostic purposes.
		bodySnippet, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
		return nil, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(bodySnippet))
	}

	// Decode the JSON payload.
	var f model.Forecast
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&f); err != nil {
		return nil, fmt.Errorf("decoding JSON: %w", err)
	}

	return &f, nil
}
