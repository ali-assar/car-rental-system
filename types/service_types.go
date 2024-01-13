package types

type OBUData struct {
	OBUID int     `json:obuID`
	Lat   float64 `json:lat`
	Long  float64 `json:long`
}

type Distance struct {
	Value float64 `json:"value"`
	OBUID int     `json:"obuID"`
	Unix  int64   `json:"unix"`
}
