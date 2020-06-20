package customer

import (
	"math"
)

// IsInvitedForBeer calculates the distance of customer from Intercom office
// and returns true if customer is in 100km radius, false if not
// This function is modified from the one at https://www.geodatasource.com/developers/go
// One point in the function is constant i.e. Intercom office
func (c Customer) IsInvitedForBeer() bool {

	const officeLat = 53.339428
	const officeLan = -6.257664
	const officeRad = float64(100)

	// Convert lattittude to radians
	radlat1 := float64(math.Pi * officeLat / 180)
	radlat2 := float64(math.Pi * c.Latt / 180)

	lng1 := float64(officeLan)
	lng2 := c.Long

	theta := float64(lng1 - lng2)
	radtheta := float64(math.Pi * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515

	dist = dist * 1.609344
	return dist <= officeRad
}
