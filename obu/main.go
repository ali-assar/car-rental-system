package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var sendInterval = time.Second

type OBUData struct {
	OBUID int     `json:obuID`
	Lat   float64 `json:lat`
	Long  float64
}

func genLatLong() (float64, float64) {
	return genCoord(), genCoord()
}

func genCoord() float64 {
	n := float64(rand.Intn(100) + 1)
	f := rand.Float64()
	return n + f

}
func main() {

	for {
		obuIDs := generateOBUIDs(20)
		for i := 0; i < len(obuIDs); i++ {
			lat, long := genLatLong()
			data := OBUData{
				OBUID: obuIDs[1],
				Lat:   lat,
				Long:  long,
			}
			fmt.Printf("%v\n", data)
		}

		time.Sleep(sendInterval)
	}
}

func generateOBUIDs(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(math.MaxInt)
	}
	return ids
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
