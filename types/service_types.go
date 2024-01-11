package types

type OBUData struct {
	OBUID int     `json:obuID`
	Lat   float64 `json:lat`
	Long  float64
}
