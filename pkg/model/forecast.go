package model

// Forecast represents the entire JSON structure from the weather API.
// The struct tags are set to match the JSON field names.
// Only the fields present in forecast.json are included.
// If additional fields are returned in the future, they can be added with the omitempty tag.

type Forecast struct {
	Latitude         float64     `json:"latitude"`
	Longitude        float64     `json:"longitude"`
	GenerationTimeMs float64     `json:"generationtime_ms"`
	UTCOffsetSeconds int         `json:"utc_offset_seconds"`
	Timezone         string      `json:"timezone"`
	TimezoneAbbrev   string      `json:"timezone_abbreviation"`
	Elevation        float64     `json:"elevation"`
	HourlyUnits      HourlyUnits `json:"hourly_units"`
	Hourly           Hourly      `json:"hourly"`
}

// HourlyUnits describes the units for each hourly field.
// The struct tags match the keys in the "hourly_units" object.
// All fields are strings as the API returns unit strings.

type HourlyUnits struct {
	Time               string `json:"time"`
	Temperature2m      string `json:"temperature_2m"`
	RelativeHumidity2m string `json:"relative_humidity_2m"`
	Rain               string `json:"rain"`
	WindSpeed10m       string `json:"wind_speed_10m"`
}

// Hourly holds the slice data for each hourly metric.
// All fields are slices of appropriate Go types.
// time is a slice of strings in ISO8601 format.
// temperature_2m, rain, wind_speed_10m are float64 slices.
// relative_humidity_2m is an int slice.

type Hourly struct {
	Time               []string  `json:"time"`
	Temperature2m      []float64 `json:"temperature_2m"`
	RelativeHumidity2m []int     `json:"relative_humidity_2m"`
	Rain               []float64 `json:"rain"`
	WindSpeed10m       []float64 `json:"wind_speed_10m"`
}
