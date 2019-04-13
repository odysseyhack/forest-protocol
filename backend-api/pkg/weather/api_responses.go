package weather

type WeatherServiceAlmanacResponse struct {
	MetaData MetaData `json:"metadata"`
	MonthData []MonthData `json:"almanac_summaries"`

}
type MetaData struct {
	Language string `json:"language"`
	TransactionID string `json:"transaction_id"`
	Version string `json:"version"`
	Latitude float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Time	int64 `json:"expire_time_gmt"`
	StatusCode int `json:"status_code"`
}

type MonthData struct {
	Class string	`json:"class"`
	StationID string `json:"stationd_id"`
	StationName string `json:"station_name"`
	Month	string `json:"almanac_dt"`
	Interval	string `json:"interval"`
	AvgHigh float32 `json:"avg_hi"`
	AvgLow float32 `json:"avg_lo"`
	RecordHigh float32 `json:"record_hi"`
	RecordHighYear int `json:"record_hi_yr"`
	RecordLow float32 `json:"record_lo"`
	RecordLowYear int `json:"record_lo_yr"`
	MeanTemp	float32 `json:"mean_temp"`
	AvgPrecip	float32 `json:"avg_precip"`
	AvgSnow	float32 `json:"avg_snow"`
	RecordPeriod int `json:"record_period"`
}
