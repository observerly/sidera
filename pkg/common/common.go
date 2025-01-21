/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/sidera
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package common

/*****************************************************************************************************************/

import (
	"math"
)

/*****************************************************************************************************************/

var DEGREES_TO_RADIANS = math.Pi / 180

/*****************************************************************************************************************/

type EclipticCoordinate struct {
	Longitude float64
	Latitude  float64
}

/*****************************************************************************************************************/

type EquatorialCoordinate struct {
	RightAscension float64
	Declination    float64
}

/*****************************************************************************************************************/

type GeographicCoordinate struct {
	Latitude  float64
	Longitude float64
	Elevation float64
}

/*****************************************************************************************************************/

type HorizontalCoordinate struct {
	Azimuth  float64
	Altitude float64
}

/*****************************************************************************************************************/

func Radians(degrees float64) float64 {
	return degrees * DEGREES_TO_RADIANS
}

/*****************************************************************************************************************/

func Degrees(radians float64) float64 {
	return radians / DEGREES_TO_RADIANS
}

/*****************************************************************************************************************/

func DegreesToHDMS(value float64) (int, int, int, float64) {
	// Calculate the degrees part:
	degrees := math.Floor(value)

	// Calculate the hours part (15 degrees per hour):
	hours := degrees / 15

	// Calculate the minutes part:
	minutes := int(math.Floor((value - degrees) * 60))

	// Calculate the seconds part:
	seconds := ((value-degrees)*60 - float64(minutes)) * 60

	// Return the HDMS values:
	return int(hours), int(degrees), int(minutes), seconds
}

/*****************************************************************************************************************/
